// The MIT License (MIT)
//
// Copyright (c) 2013-2017 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// The oryx AVC package includes some utilites.
// The NALU(Netowrk Abstraction Layer Unit) is suitable for transmission over network.
// We could package NALUs by AnnexB, IBMF or RTP according to different scenarios.
// 	@note AnnexB is designed for bit-oriented stream, such as MPEG-TS/HLS, please
// 		read ISO_IEC_14496-10-AVC-2003.pdf at page 211, AnnexB Byte stream Format.
//	@note IBMF is designed for file storage, such as MP4/FLV, please read
//		ISO_IEC_14496-15-AVC-format-2012.pdf at page 16, 5.2.4.1 AVC decoder
// 		configuration record.
//	@note RTP payload for H.264, defined in RFC6184 https://tools.ietf.org/html/rfc6184
//		it directly uses and extends the NAL header.
package avc

import (
	"github.com/ossrs/go-oryx-lib/errors"
)

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 44, 7.3.1 NAL unit syntax
type NALRefIDC uint8

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 44, 7.3.1 NAL unit syntax
type NALUType uint8

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 60, 7.4.1 NAL unit semantics
type NALUHeader struct {
	// The 2-bits nal_ref_idc.
	nalRefIDC NALRefIDC
	// The 5-bits nal_unit_type.
	naluType NALUType
}

func newNALUHeader() *NALUHeader {
	return &NALUHeader{}
}

// @doc ISO_IEC_14496-10-AVC-2003.pdf at page 60, 7.4.1 NAL unit semantics
type NALU struct {
	*NALUHeader
	data []byte
}

func newNALU() *NALU {
	return &NALU{NALUHeader: newNALUHeader()}
}

func (v *NALU) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return errors.New("empty NALU")
	}
	v.nalRefIDC = NALRefIDC(uint8(data[0]>>5) & 0x03)
	v.naluType = NALUType(uint8(data[0]) & 0x1f)
	v.data = data[1:]
	return nil
}

func (v *NALU) MarshalBinary() ([]byte, error) {
	h := []byte{
		byte(v.nalRefIDC)<<5 | byte(v.naluType),
	}
	if len(v.data) == 0 {
		return h, nil
	}
	return append(h, v.data...), nil
}

// @doc Annex A Profiles and levels, ISO_IEC_14496-10-AVC-2003.pdf, page 205.
// @note that it's uint8 in IBMF, but extended in other specs, so we use uint16.
type AVCProfile uint16

const (
	// @see ffmpeg, libavcodec/avcodec.h:2713
	AVCProfileBaseline            AVCProfile = 66
	AVCProfileConstrainedBaseline AVCProfile = 578
	AVCProfileMain                AVCProfile = 77
	AVCProfileExtended            AVCProfile = 88
	AVCProfileHigh                AVCProfile = 100
	AVCProfileHigh10              AVCProfile = 110
	AVCProfileHigh10Intra         AVCProfile = 2158
	AVCProfileHigh422             AVCProfile = 122
	AVCProfileHigh422Intra        AVCProfile = 2170
	AVCProfileHigh444             AVCProfile = 144
	AVCProfileHigh444Predictive   AVCProfile = 244
	AVCProfileHigh444Intra        AVCProfile = 2192
)

func (v AVCProfile) String() string {
	switch v {
	case AVCProfileBaseline:
		return "Baseline"
	case AVCProfileConstrainedBaseline:
		return "ConstrainedBaseline"
	case AVCProfileMain:
		return "Main"
	case AVCProfileExtended:
		return "Extended"
	case AVCProfileHigh:
		return "High"
	case AVCProfileHigh10:
		return "High10"
	case AVCProfileHigh10Intra:
		return "High10Intra"
	case AVCProfileHigh422:
		return "High422"
	case AVCProfileHigh422Intra:
		return "High422Intra"
	case AVCProfileHigh444:
		return "High444"
	case AVCProfileHigh444Predictive:
		return "High444Predictive"
	case AVCProfileHigh444Intra:
		return "High444Intra"
	default:
		return "Forbidden"
	}
}

// @doc Annex A Profiles and levels, ISO_IEC_14496-10-AVC-2003.pdf, page 207.
type AVCLevel uint8

const (
	AVCLevel_1  = 10
	AVCLevel_11 = 11
	AVCLevel_12 = 12
	AVCLevel_13 = 13
	AVCLevel_2  = 20
	AVCLevel_21 = 21
	AVCLevel_22 = 22
	AVCLevel_3  = 30
	AVCLevel_31 = 31
	AVCLevel_32 = 32
	AVCLevel_4  = 40
	AVCLevel_41 = 41
	AVCLevel_5  = 50
	AVCLevel_51 = 51
)

func (v AVCLevel) String() string {
	switch v {
	case AVCLevel_1:
		return "Level_1"
	case AVCLevel_11:
		return "Level_11"
	case AVCLevel_12:
		return "Level_12"
	case AVCLevel_13:
		return "Level_13"
	case AVCLevel_2:
		return "Level_2"
	case AVCLevel_21:
		return "Level_21"
	case AVCLevel_22:
		return "Level_22"
	case AVCLevel_3:
		return "Level_3"
	case AVCLevel_31:
		return "Level_31"
	case AVCLevel_32:
		return "Level_32"
	case AVCLevel_4:
		return "Level_4"
	case AVCLevel_41:
		return "Level_41"
	case AVCLevel_5:
		return "Level_5"
	case AVCLevel_51:
		return "Level_51"
	default:
		return "Forbidden"
	}
}

// @doc ISO_IEC_14496-15-AVC-format-2012.pdf at page 16, 5.2.4.1.1 Syntax
type AVCDecoderConfigurationRecord struct {
	// It contains the profile code as defined in ISO/IEC 14496-10.
	configurationVersion uint8
	// It is a byte defined exactly the same as the byte which occurs between the
	// profile_IDC and level_IDC in a sequence parameter set (SPS), as defined in
	// ISO/IEC 14496-10.
	// @remark It's 8 bits.
	avcProfileIndication AVCProfile
	// It contains the level code as defined in ISO/IEC 14496-10.
	profileCompatibility uint8
	// It indicates the length in bytes of the NALUnitLength field in an AVC video sample
	// or AVC parameter set sample of the associated stream minus one.
	avcLevelIndication AVCLevel
	// It indicates the length in bytes of the NALUnitLength field in an AVC video sample
	// or AVC parameter set sample of the associated stream minus one.
	lengthSizeMinusOne uint8
	// It contains a SPS NAL unit, as specified in ISO/IEC 14496-10. SPSs shall occur in
	// order of ascending parameter set identifier with gaps being allowed.
	sequenceParameterSetNALUnits []*NALU
	// It contains a PPS NAL unit, as specified in ISO/IEC 14496-10. PPSs shall occur in
	// order of ascending parameter set identifier with gaps being allowed.
	pictureParameterSetNALUnits []*NALU
	// @remark We ignore the sequenceParameterSetExtNALUnit.
}

func NewAVCDecoderConfigurationRecord() *AVCDecoderConfigurationRecord {
	return &AVCDecoderConfigurationRecord{}
}

func (v *AVCDecoderConfigurationRecord) UnmarshalBinary(data []byte) error {
	b := data
	if len(b) < 6 {
		return errors.Errorf("requires 6+ only %v bytes", len(b))
	}

	v.configurationVersion = uint8(b[0])
	v.avcProfileIndication = AVCProfile(uint8(b[1]))
	v.profileCompatibility = uint8(b[2])
	v.avcLevelIndication = AVCLevel(uint8(b[3]))
	v.lengthSizeMinusOne = uint8(b[4]) & 0x03
	b = b[4:]

	numOfSequenceParameterSets := uint8(b[0]) & 0x1f
	b = b[1:]
	for i := 0; i < numOfSequenceParameterSets; i++ {
		if len(b) < 2 {
			return errors.Errorf("requires 2+ only %v bytes", len(b))
		}
		sequenceParameterSetLength := int(uint16(b[0])<<8 | uint16(b[1]))
		b = b[2:]

		if len(b) < sequenceParameterSetLength {
			return errors.Errorf("requires %v only %v bytes", sequenceParameterSetLength, len(b))
		}
		sps := newNALU()
		if err := sps.UnmarshalBinary(b[:sequenceParameterSetLength]); err != nil {
			return errors.WithMessage(err, "unmarshal")
		}
		b = b[sequenceParameterSetLength:]

		v.sequenceParameterSetNALUnits = append(v.sequenceParameterSetNALUnits, sps)
	}

	if len(b) < 1 {
		return errors.New("no PPS length")
	}
	numOfPictureParameterSets := uint8(b[0])
	b = b[1:]
	for i := 0; i < numOfPictureParameterSets; i++ {
		if len(b) < 2 {
			return errors.Errorf("requiers 2+ only %v bytes", len(b))
		}

		pictureParameterSetLength := int(uint16(b[0])<<8 | uint16(b[1]))
		b = b[2:]

		if len(b) < pictureParameterSetLength {
			return errors.Errorf("requires %v only %v bytes", pictureParameterSetLength, len(b))
		}
		pps := newNALU()
		if err := pps.UnmarshalBinary(b[:pictureParameterSetLength]); err != nil {
			return errors.WithMessage(err, "unmarshal")
		}
		b = b[pictureParameterSetLength:]

		v.pictureParameterSetNALUnits = append(v.pictureParameterSetNALUnits, pps)
	}
	return nil
}
