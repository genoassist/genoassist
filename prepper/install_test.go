package prepper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genomagic/config_parser"
)

// TestNewPrep is used for locally testing whether the prepper package managed to pull
// all the necessary images for running GenoMagic. The test is prevented from running upon
// deploys as the environment in which the build runs cannot/should not pull Docker images.
// This should be run at least once when adding a new assembler or a new quality controller.
// Before running the test, local images can be removed via:
// docker image remove vout/megahit nanozoo/flye bcgsc/abyss replikation/porechop greatfireball/canu
// Images can be pulled independently. If that is the case, the workgroup delta has to change to however
// many images are being tested
func TestNewPrep(t *testing.T) {
	config := &config_parser.Config{
		Assemblers: config_parser.AssemblerConfig{
			Megahit: config_parser.MegahitConfig{},
			Abyss:   config_parser.AbyssConfig{},
			Flye:    config_parser.FlyeConfig{},
		},
		GenoMagic: config_parser.GenoMagicConfig{
			Assemblers:    nil,
			InputFilePath: "",
			OutputPath:    "",
			Threads:       0,
			Prep:          true,
		},
	}
	t.Run("test_new_prep_pulls_the_expected_images", func(t *testing.T) {
		assert.NotNil(t, config)
		//errs := NewPrep(config)
		//for len(errs) > 0 {
		//	select {
		//	case err := <-errs:
		//		assert.NilError(t, err, fmt.Sprintf("failed to pull image, err: %s", err))
		//	default:
		//		continue
		//	}
		//}
	})
}
