// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package nvml

// #include <sys/prctl.h>
import "C"

// nvml.SystemGetDriverVersion()
func SystemGetDriverVersion() (string, Return) {
	Version := make([]byte, SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	ret := nvmlSystemGetDriverVersion(&Version[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.SystemGetNVMLVersion()
func SystemGetNVMLVersion() (string, Return) {
	Version := make([]byte, SYSTEM_NVML_VERSION_BUFFER_SIZE)
	ret := nvmlSystemGetNVMLVersion(&Version[0], SYSTEM_NVML_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// nvml.SystemGetCudaDriverVersion()
func SystemGetCudaDriverVersion() (int, Return) {
	var CudaDriverVersion int32
	ret := nvmlSystemGetCudaDriverVersion(&CudaDriverVersion)
	return int(CudaDriverVersion), ret
}

// nvml.SystemGetCudaDriverVersion_v2()
func SystemGetCudaDriverVersion_v2() (int, Return) {
	var CudaDriverVersion int32
	ret := nvmlSystemGetCudaDriverVersion_v2(&CudaDriverVersion)
	return int(CudaDriverVersion), ret
}

// nvml.SystemGetProcessName()
func SystemGetProcessName(Pid int) (string, Return) {
	Name := make([]byte, C.PR_SET_NAME)
	ret := nvmlSystemGetProcessName(uint32(Pid), &Name[0], C.PR_SET_NAME)
	return string(Name[:clen(Name)]), ret
}

// nvml.SystemGetHicVersion()
func SystemGetHicVersion() ([]HwbcEntry, Return) {
	var HwbcCount uint32 = 1 // Will be reduced upon returning
	for {
		HwbcEntries := make([]HwbcEntry, HwbcCount)
		ret := nvmlSystemGetHicVersion(&HwbcCount, &HwbcEntries[0])
		if ret == SUCCESS {
			return HwbcEntries[:HwbcCount], ret
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		HwbcCount *= 2
	}
}

// nvml.SystemGetTopologyGpuSet()
func SystemGetTopologyGpuSet(CpuNumber int) ([]Device, Return) {
	var Count uint32
	ret := nvmlSystemGetTopologyGpuSet(uint32(CpuNumber), &Count, nil)
	if ret != SUCCESS {
		return nil, ret
	}
	if Count == 0 {
		return []Device{}, ret
	}
	DeviceArray := make([]Device, Count)
	ret = nvmlSystemGetTopologyGpuSet(uint32(CpuNumber), &Count, &DeviceArray[0])
	return DeviceArray, ret
}