package docker_applications_test

import (
	"github.com/Lajule/docker-applications"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	config := docker_applications.Config{
		Version: "1",
		Applications: map[string]docker_applications.Application{
			"front": docker_applications.Application{
				Dir: "dir",
				DependsOn: []string{
					"api",
					"varnish",
				},
			},
			"api": docker_applications.Application{
				Dir:  "dir",
				File: "file",
			},
			"varnish": docker_applications.Application{
				Dir: "dir",
			},
			"circular": docker_applications.Application{
				Dir: "dir",
				DependsOn: []string{
					"circular",
				},
			},
			"nodir": docker_applications.Application{},
		},
	}

	tests := []struct {
		name        string
		application string
		want        []string
	}{
		{
			name:        "Basic dependency",
			application: "front",
			want:        []string{"-f", "dir/docker-compose.yml", "-f", "dir/file", "-f", "dir/docker-compose.yml", "up"},
		},
		{
			name:        "Circular dependency",
			application: "circular",
			want:        []string{"-f", "dir/docker-compose.yml", "up"},
		},
		{
			name:        "Undefined Application",
			application: "undefined",
			want:        nil,
		},
		{
			name:        "Undefined directory in Application",
			application: "nodir",
			want:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := config.Parse(tc.application, []string{"up"})
			if tc.want == nil {
				if err == nil {
					t.Errorf("Parse() = (%#v, %v), want (nil, err)", got, err)
				}
			} else {
				if err != nil || !reflect.DeepEqual(got, tc.want) {
					t.Errorf("Parse() = (%#v, %v), want (%#v, nil)", got, err, tc.want)
				}
			}
		})
	}
}
