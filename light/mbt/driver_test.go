package mbt

import (
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/light"
	"github.com/tendermint/tendermint/types"
)

const jsonDir = "./json"

func TestVerify(t *testing.T) {
	filenames := jsonFilenames(t)

	for _, filename := range filenames {
		t.Logf("-> %s", filename)

		jsonBlob, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatal(err)
		}

		var tc testCase
		err = tmjson.Unmarshal(jsonBlob, &tc)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(tc.Description)

		var (
			trustedSignedHeader = tc.Initial.SignedHeader
			trustedNextVals     = tc.Initial.NextValidatorSet
			trustingPeriod      = time.Duration(tc.Initial.TrustingPeriod) * time.Microsecond
		)

		for _, input := range tc.Input {
			var (
				newSignedHeader = input.LightBlock.SignedHeader
				newVals         = input.LightBlock.ValidatorSet
			)

			err = light.Verify(
				&trustedSignedHeader,
				&trustedNextVals,
				newSignedHeader,
				newVals,
				trustingPeriod,
				input.Now,
				1*time.Second,
				light.DefaultTrustLevel,
			)
			if err != nil {
				t.Logf("TRUSTED HEADER %v\nNEW HEADER %v", trustedSignedHeader, newSignedHeader)

				switch input.Verdict {
				case "SUCCESS":
					t.Fatalf("unexpected error: %v", err)
				case "NOT_ENOUGH_TRUST":
					if _, ok := err.(light.ErrNewValSetCantBeTrusted); !ok {
						t.Fatalf("expected ErrNewValSetCantBeTrusted, but got %v", err)
					}
				case "INVALID":
					if _, ok := err.(light.ErrInvalidHeader); !ok {
						t.Fatalf("expected ErrInvalidHeader, but got %v", err)
					}
				default:
					t.Fatalf("unexpected verdict: %q", input.Verdict)
				}
			} else { // advance
				trustedSignedHeader = *newSignedHeader
				trustedNextVals = *input.LightBlock.NextValidatorSet
			}
		}
	}
}

// jsonFilenames returns a list of files in jsonDir directory
func jsonFilenames(t *testing.T) []string {
	names := make([]string, 0)

	files, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		t.Fatal(err)
	}

	base := filepath.Base(jsonDir)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			names = append(names, filepath.Join(base, file.Name()))
		}
	}

	return names
}

type testCase struct {
	Description string      `json:"description"`
	Initial     initialData `json:"initial"`
	Input       []inputData `json:"input"`
}

type initialData struct {
	SignedHeader     types.SignedHeader `json:"signed_header"`
	NextValidatorSet types.ValidatorSet `json:"next_validator_set"`
	TrustingPeriod   uint64             `json:"trusting_period"`
	Now              time.Time          `json:"now"`
}

type inputData struct {
	LightBlock lightBlockWithNextValidatorSet `json:"block"`
	Now        time.Time                      `json:"now"`
	Verdict    string                         `json:"verdict"`
}

// In tendermint-rs, NextValidatorSet is used to verify new blocks (opposite to
// Go tendermint).
type lightBlockWithNextValidatorSet struct {
	*types.SignedHeader `json:"signed_header"`
	ValidatorSet        *types.ValidatorSet `json:"validator_set"`
	NextValidatorSet    *types.ValidatorSet `json:"next_validator_set"`
}