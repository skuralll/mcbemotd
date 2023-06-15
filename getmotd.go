package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sandertv/go-raknet"
)

// do not split if \ exists before separator
func customSplit(text string, separator rune) []string {
	var result []string
	var sb strings.Builder
	escaped := false

	for _, c := range text {
		if c == '\\' && !escaped {
			escaped = true
		} else if c == separator && !escaped {
			result = append(result, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(c)
			escaped = false
		}
	}

	result = append(result, sb.String())

	return result
}

type ServerInfo struct {
	Edition    string // Game edition, MCPE or MCEE(Education Edition)
	Motd1      string // MOTD line 1
	Protocol   int    // Protocol version
	Version    string // Game version name
	Players    int    // Player Count
	PlayersMax int    // Max Player Count
	Uid        string // Server unique id
	Motd2      string // MOTD line 2
	ModeStr    string // Game mode (string)
	ModeNum    int    // Game mode (numeric)
	Portv4     int    // Srever port (v4)
	Portv6     int    // Srever port (v6)
}

// decode from bytes
func (st ServerInfo) Decode(b []byte) (ServerInfo, error) {
	// split by ;
	infoArr := customSplit(string(b), ';')
	// assignments
	if len(infoArr) < 10 {
		return st, fmt.Errorf("byte array has too few elements")
	}
	st.Edition = infoArr[0]
	st.Motd1 = infoArr[1]
	st.Protocol, _ = strconv.Atoi(infoArr[2])
	st.Version = infoArr[3]
	st.Players, _ = strconv.Atoi(infoArr[4])
	st.PlayersMax, _ = strconv.Atoi(infoArr[5])
	st.Uid = infoArr[6]
	st.Motd2 = infoArr[7]
	st.ModeStr = infoArr[8]
	st.ModeNum, _ = strconv.Atoi(infoArr[9])
	// May not exist below
	if len(infoArr) < 11 {
		st.Portv4 = -1
	} else {
		st.Portv4, _ = strconv.Atoi(infoArr[10])
	}
	if len(infoArr) < 12 {
		st.Portv6 = -1
	} else {
		st.Portv6, _ = strconv.Atoi(infoArr[11])
	}
	return st, nil
}

// decode too
func Decode(b []byte) (ServerInfo, error) {
	return ServerInfo{}.Decode(b)
}

// get server info from address (ex: 0.0.0.0:19132)
func GetServerInfo(address string) (ServerInfo, error) {
	// get bytes
	b, err := raknet.Ping(address)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("ping failed %w", err)
	}
	// decode
	info, err := Decode(b)
	if err != nil {

		return ServerInfo{}, fmt.Errorf("decode failed: %w", err)
	}
	return info, nil
}
