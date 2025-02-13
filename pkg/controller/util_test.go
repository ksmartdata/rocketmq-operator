package controller

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func getTZ(envs []corev1.EnvVar) string {
	for _, item := range envs {
		if item.Name == "TZ" {
			return item.Value
		}
	}
	return ""
}
func TestSetDefaultTZ(t *testing.T) {
	type args struct {
		envs []corev1.EnvVar
		tz   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no TZ",
			args: args{
				envs: []corev1.EnvVar{
					{
						Name:  "name1",
						Value: "value1",
					},
				},
				tz: "Asia/Bangkok",
			},
			want: "Asia/Bangkok",
		},
		{
			name: "have TZ",
			args: args{
				envs: []corev1.EnvVar{
					{
						Name:  "TZ",
						Value: "Asia/Shanghai",
					},
				},
				tz: "Asia/Bangkok",
			},
			want: "Asia/Shanghai",
		},
		{
			name: "have TZ, empty value",
			args: args{
				envs: []corev1.EnvVar{
					{
						Name:  "TZ",
						Value: "",
					},
				},
				tz: "Asia/Bangkok",
			},
			want: "Asia/Bangkok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envs := SetDefaultTZ(tt.args.envs, tt.args.tz)
			got := getTZ(envs)
			if got != tt.want {
				t.Fatalf("want: %s, got: %s", tt.want, got)
			}
		})
	}
}
