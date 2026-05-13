/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package tun

import (
	"errors"
	"os"
	"sync/atomic"

	"golang.org/x/sys/windows"
)

var errWintunDisabled = errors.New("wintun support disabled")

type NativeTun struct {
	rate   rateJuggler
	events chan Event
}

type rateJuggler struct {
	current       atomic.Uint64
	nextByteCount atomic.Uint64
	nextStartTime atomic.Int64
	changing      atomic.Bool
}

var (
	WintunTunnelType          = "WireGuard"
	WintunStaticRequestedGUID *windows.GUID
)

// CreateTUN would normally create a Wintun interface. This fork is used by
// mihomo's tsnet integration, which runs entirely in userspace and must not
// create a system interface.
func CreateTUN(ifname string, mtu int) (Device, error) {
	return nil, errWintunDisabled
}

func CreateTUNWithRequestedGUID(ifname string, requestedGUID *windows.GUID, mtu int) (Device, error) {
	return nil, errWintunDisabled
}

func (tun *NativeTun) Name() (string, error) {
	return "", errWintunDisabled
}

func (tun *NativeTun) File() *os.File {
	return nil
}

func (tun *NativeTun) Events() <-chan Event {
	if tun.events == nil {
		tun.events = make(chan Event)
		close(tun.events)
	}
	return tun.events
}

func (tun *NativeTun) Close() error {
	return nil
}

func (tun *NativeTun) MTU() (int, error) {
	return 0, errWintunDisabled
}

func (tun *NativeTun) ForceMTU(mtu int) {}

func (tun *NativeTun) BatchSize() int {
	return 1
}

func (tun *NativeTun) Read(bufs [][]byte, sizes []int, offset int) (int, error) {
	return 0, errWintunDisabled
}

func (tun *NativeTun) Write(bufs [][]byte, offset int) (int, error) {
	return 0, errWintunDisabled
}

func (tun *NativeTun) LUID() uint64 {
	return 0
}

func (tun *NativeTun) RunningVersion() (version uint32, err error) {
	return 0, errWintunDisabled
}
