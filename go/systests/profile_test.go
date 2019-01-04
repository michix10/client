package systests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProofSuggestions(t *testing.T) {
	tt := newTeamTester(t)
	defer tt.cleanup()

	alice := tt.addUser("abc")

	profileProofs, err := alice.userClient.ProfileProofSuggestions(context.Background(), 0)
	require.NoError(t, err)
	t.Logf("%v", profileProofs)

	proofs, err := alice.userClient.ProofSuggestions(context.Background(), 0)
	require.NoError(t, err)
	t.Logf("%v", proofs)
}
