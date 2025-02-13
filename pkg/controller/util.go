package controller

import corev1 "k8s.io/api/core/v1"

func SetDefaultTZ(envs []corev1.EnvVar, tz string) []corev1.EnvVar {
	for index, item := range envs {
		if item.Name == "TZ" {
			if item.Value == "" {
				envs[index].Value = tz
				return envs
			}
			return envs
		}
	}
	return append(envs, corev1.EnvVar{Name: "TZ", Value: tz})
}
