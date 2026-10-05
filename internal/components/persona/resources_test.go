package persona_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jonsanchezr/agent-smith/v4/internal/components/persona"
	"github.com/jonsanchezr/agent-smith/v4/internal/model"
)

func TestResourcePlanOutputStylePaths(t *testing.T) {
	dir := t.TempDir()
	agentSmith := filepath.Join(dir, "agent-smith.md")
	neutral := filepath.Join(dir, "neutral.md")

	tests := []struct {
		name    string
		persona model.PersonaID
		want    persona.OutputStylePaths
	}{
		{
			name:    "agent-smith writes its selected style without removing neutral",
			persona: model.PersonaAgentSmith,
			want: persona.OutputStylePaths{
				Write:  agentSmith,
				Backup: []string{agentSmith, neutral},
			},
		},
		{
			name:    "neutral writes its selected style and removes retired agent-smith",
			persona: model.PersonaNeutral,
			want: persona.OutputStylePaths{
				Write:  neutral,
				Backup: []string{agentSmith, neutral},
				Remove: []string{agentSmith},
			},
		},
		{
			name:    "legacy neutral alias writes neutral and removes retired agent-smith",
			persona: model.PersonaAgentSmithNeutralArtifacts,
			want: persona.OutputStylePaths{
				Write:  neutral,
				Backup: []string{agentSmith, neutral},
				Remove: []string{agentSmith},
			},
		},
		{
			name:    "custom manages no output styles",
			persona: model.PersonaCustom,
			want:    persona.OutputStylePaths{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := persona.ResourcePlanFor(tt.persona).OutputStylePaths(dir)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ResourcePlanFor(%q).OutputStylePaths() = %#v, want %#v", tt.persona, got, tt.want)
			}
		})
	}
}

