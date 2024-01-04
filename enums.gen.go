package ffmpeg

// #include <libavcodec/avcodec.h>
// #include <libavcodec/codec.h>
// #include <libavcodec/codec_desc.h>
// #include <libavcodec/codec_id.h>
// #include <libavcodec/codec_par.h>
// #include <libavcodec/defs.h>
// #include <libavcodec/packet.h>
// #include <libavfilter/avfilter.h>
// #include <libavfilter/buffersink.h>
// #include <libavfilter/buffersrc.h>
// #include <libavformat/avformat.h>
// #include <libavformat/avio.h>
// #include <libavutil/avutil.h>
// #include <libavutil/buffer.h>
// #include <libavutil/channel_layout.h>
// #include <libavutil/dict.h>
// #include <libavutil/frame.h>
// #include <libavutil/hwcontext.h>
// #include <libavutil/log.h>
// #include <libavutil/mem.h>
// #include <libavutil/opt.h>
// #include <libavutil/pixfmt.h>
// #include <libavutil/rational.h>
// #include <libavutil/samplefmt.h>
import "C"

// --- Enum AVSubtitleType ---

// AVSubtitleType wraps AVSubtitleType.
type AVSubtitleType C.enum_AVSubtitleType

const (
	// SubtitleNone wraps SUBTITLE_NONE.
	SubtitleNone AVSubtitleType = C.SUBTITLE_NONE
	// SubtitleBitmap wraps SUBTITLE_BITMAP.
	SubtitleBitmap AVSubtitleType = C.SUBTITLE_BITMAP
	// SubtitleText wraps SUBTITLE_TEXT.
	SubtitleText AVSubtitleType = C.SUBTITLE_TEXT
	// SubtitleAss wraps SUBTITLE_ASS.
	SubtitleAss AVSubtitleType = C.SUBTITLE_ASS
)

// --- Enum AVPictureStructure ---

// AVPictureStructure wraps AVPictureStructure.
type AVPictureStructure C.enum_AVPictureStructure

const (
	// AVPictureStructureUnknown wraps AV_PICTURE_STRUCTURE_UNKNOWN.
	AVPictureStructureUnknown AVPictureStructure = C.AV_PICTURE_STRUCTURE_UNKNOWN
	// AVPictureStructureTopField wraps AV_PICTURE_STRUCTURE_TOP_FIELD.
	AVPictureStructureTopField AVPictureStructure = C.AV_PICTURE_STRUCTURE_TOP_FIELD
	// AVPictureStructureBottomField wraps AV_PICTURE_STRUCTURE_BOTTOM_FIELD.
	AVPictureStructureBottomField AVPictureStructure = C.AV_PICTURE_STRUCTURE_BOTTOM_FIELD
	// AVPictureStructureFrame wraps AV_PICTURE_STRUCTURE_FRAME.
	AVPictureStructureFrame AVPictureStructure = C.AV_PICTURE_STRUCTURE_FRAME
)

// --- Enum AVCodecID ---

// AVCodecID wraps AVCodecID.
type AVCodecID C.enum_AVCodecID

const (
	// AVCodecIdNone wraps AV_CODEC_ID_NONE.
	AVCodecIdNone AVCodecID = C.AV_CODEC_ID_NONE
	// AVCodecIdMpeg1Video wraps AV_CODEC_ID_MPEG1VIDEO.
	AVCodecIdMpeg1Video AVCodecID = C.AV_CODEC_ID_MPEG1VIDEO
	// AVCodecIdMpeg2Video wraps AV_CODEC_ID_MPEG2VIDEO.
	AVCodecIdMpeg2Video AVCodecID = C.AV_CODEC_ID_MPEG2VIDEO
	// AVCodecIdH261 wraps AV_CODEC_ID_H261.
	AVCodecIdH261 AVCodecID = C.AV_CODEC_ID_H261
	// AVCodecIdH263 wraps AV_CODEC_ID_H263.
	AVCodecIdH263 AVCodecID = C.AV_CODEC_ID_H263
	// AVCodecIdRv10 wraps AV_CODEC_ID_RV10.
	AVCodecIdRv10 AVCodecID = C.AV_CODEC_ID_RV10
	// AVCodecIdRv20 wraps AV_CODEC_ID_RV20.
	AVCodecIdRv20 AVCodecID = C.AV_CODEC_ID_RV20
	// AVCodecIdMjpeg wraps AV_CODEC_ID_MJPEG.
	AVCodecIdMjpeg AVCodecID = C.AV_CODEC_ID_MJPEG
	// AVCodecIdMjpegb wraps AV_CODEC_ID_MJPEGB.
	AVCodecIdMjpegb AVCodecID = C.AV_CODEC_ID_MJPEGB
	// AVCodecIdLjpeg wraps AV_CODEC_ID_LJPEG.
	AVCodecIdLjpeg AVCodecID = C.AV_CODEC_ID_LJPEG
	// AVCodecIdSp5X wraps AV_CODEC_ID_SP5X.
	AVCodecIdSp5X AVCodecID = C.AV_CODEC_ID_SP5X
	// AVCodecIdJpegls wraps AV_CODEC_ID_JPEGLS.
	AVCodecIdJpegls AVCodecID = C.AV_CODEC_ID_JPEGLS
	// AVCodecIdMpeg4 wraps AV_CODEC_ID_MPEG4.
	AVCodecIdMpeg4 AVCodecID = C.AV_CODEC_ID_MPEG4
	// AVCodecIdRawvideo wraps AV_CODEC_ID_RAWVIDEO.
	AVCodecIdRawvideo AVCodecID = C.AV_CODEC_ID_RAWVIDEO
	// AVCodecIdMsmpeg4V1 wraps AV_CODEC_ID_MSMPEG4V1.
	AVCodecIdMsmpeg4V1 AVCodecID = C.AV_CODEC_ID_MSMPEG4V1
	// AVCodecIdMsmpeg4V2 wraps AV_CODEC_ID_MSMPEG4V2.
	AVCodecIdMsmpeg4V2 AVCodecID = C.AV_CODEC_ID_MSMPEG4V2
	// AVCodecIdMsmpeg4V3 wraps AV_CODEC_ID_MSMPEG4V3.
	AVCodecIdMsmpeg4V3 AVCodecID = C.AV_CODEC_ID_MSMPEG4V3
	// AVCodecIdWmv1 wraps AV_CODEC_ID_WMV1.
	AVCodecIdWmv1 AVCodecID = C.AV_CODEC_ID_WMV1
	// AVCodecIdWmv2 wraps AV_CODEC_ID_WMV2.
	AVCodecIdWmv2 AVCodecID = C.AV_CODEC_ID_WMV2
	// AVCodecIdH263P wraps AV_CODEC_ID_H263P.
	AVCodecIdH263P AVCodecID = C.AV_CODEC_ID_H263P
	// AVCodecIdH263I wraps AV_CODEC_ID_H263I.
	AVCodecIdH263I AVCodecID = C.AV_CODEC_ID_H263I
	// AVCodecIdFlv1 wraps AV_CODEC_ID_FLV1.
	AVCodecIdFlv1 AVCodecID = C.AV_CODEC_ID_FLV1
	// AVCodecIdSvq1 wraps AV_CODEC_ID_SVQ1.
	AVCodecIdSvq1 AVCodecID = C.AV_CODEC_ID_SVQ1
	// AVCodecIdSvq3 wraps AV_CODEC_ID_SVQ3.
	AVCodecIdSvq3 AVCodecID = C.AV_CODEC_ID_SVQ3
	// AVCodecIdDvvideo wraps AV_CODEC_ID_DVVIDEO.
	AVCodecIdDvvideo AVCodecID = C.AV_CODEC_ID_DVVIDEO
	// AVCodecIdHuffyuv wraps AV_CODEC_ID_HUFFYUV.
	AVCodecIdHuffyuv AVCodecID = C.AV_CODEC_ID_HUFFYUV
	// AVCodecIdCyuv wraps AV_CODEC_ID_CYUV.
	AVCodecIdCyuv AVCodecID = C.AV_CODEC_ID_CYUV
	// AVCodecIdH264 wraps AV_CODEC_ID_H264.
	AVCodecIdH264 AVCodecID = C.AV_CODEC_ID_H264
	// AVCodecIdIndeo3 wraps AV_CODEC_ID_INDEO3.
	AVCodecIdIndeo3 AVCodecID = C.AV_CODEC_ID_INDEO3
	// AVCodecIdVp3 wraps AV_CODEC_ID_VP3.
	AVCodecIdVp3 AVCodecID = C.AV_CODEC_ID_VP3
	// AVCodecIdTheora wraps AV_CODEC_ID_THEORA.
	AVCodecIdTheora AVCodecID = C.AV_CODEC_ID_THEORA
	// AVCodecIdAsv1 wraps AV_CODEC_ID_ASV1.
	AVCodecIdAsv1 AVCodecID = C.AV_CODEC_ID_ASV1
	// AVCodecIdAsv2 wraps AV_CODEC_ID_ASV2.
	AVCodecIdAsv2 AVCodecID = C.AV_CODEC_ID_ASV2
	// AVCodecIdFfv1 wraps AV_CODEC_ID_FFV1.
	AVCodecIdFfv1 AVCodecID = C.AV_CODEC_ID_FFV1
	// AVCodecId4Xm wraps AV_CODEC_ID_4XM.
	AVCodecId4Xm AVCodecID = C.AV_CODEC_ID_4XM
	// AVCodecIdVcr1 wraps AV_CODEC_ID_VCR1.
	AVCodecIdVcr1 AVCodecID = C.AV_CODEC_ID_VCR1
	// AVCodecIdCljr wraps AV_CODEC_ID_CLJR.
	AVCodecIdCljr AVCodecID = C.AV_CODEC_ID_CLJR
	// AVCodecIdMdec wraps AV_CODEC_ID_MDEC.
	AVCodecIdMdec AVCodecID = C.AV_CODEC_ID_MDEC
	// AVCodecIdRoq wraps AV_CODEC_ID_ROQ.
	AVCodecIdRoq AVCodecID = C.AV_CODEC_ID_ROQ
	// AVCodecIdInterplayVideo wraps AV_CODEC_ID_INTERPLAY_VIDEO.
	AVCodecIdInterplayVideo AVCodecID = C.AV_CODEC_ID_INTERPLAY_VIDEO
	// AVCodecIdXanWc3 wraps AV_CODEC_ID_XAN_WC3.
	AVCodecIdXanWc3 AVCodecID = C.AV_CODEC_ID_XAN_WC3
	// AVCodecIdXanWc4 wraps AV_CODEC_ID_XAN_WC4.
	AVCodecIdXanWc4 AVCodecID = C.AV_CODEC_ID_XAN_WC4
	// AVCodecIdRpza wraps AV_CODEC_ID_RPZA.
	AVCodecIdRpza AVCodecID = C.AV_CODEC_ID_RPZA
	// AVCodecIdCinepak wraps AV_CODEC_ID_CINEPAK.
	AVCodecIdCinepak AVCodecID = C.AV_CODEC_ID_CINEPAK
	// AVCodecIdWsVqa wraps AV_CODEC_ID_WS_VQA.
	AVCodecIdWsVqa AVCodecID = C.AV_CODEC_ID_WS_VQA
	// AVCodecIdMsrle wraps AV_CODEC_ID_MSRLE.
	AVCodecIdMsrle AVCodecID = C.AV_CODEC_ID_MSRLE
	// AVCodecIdMsvideo1 wraps AV_CODEC_ID_MSVIDEO1.
	AVCodecIdMsvideo1 AVCodecID = C.AV_CODEC_ID_MSVIDEO1
	// AVCodecIdIdcin wraps AV_CODEC_ID_IDCIN.
	AVCodecIdIdcin AVCodecID = C.AV_CODEC_ID_IDCIN
	// AVCodecId8Bps wraps AV_CODEC_ID_8BPS.
	AVCodecId8Bps AVCodecID = C.AV_CODEC_ID_8BPS
	// AVCodecIdSmc wraps AV_CODEC_ID_SMC.
	AVCodecIdSmc AVCodecID = C.AV_CODEC_ID_SMC
	// AVCodecIdFlic wraps AV_CODEC_ID_FLIC.
	AVCodecIdFlic AVCodecID = C.AV_CODEC_ID_FLIC
	// AVCodecIdTruemotion1 wraps AV_CODEC_ID_TRUEMOTION1.
	AVCodecIdTruemotion1 AVCodecID = C.AV_CODEC_ID_TRUEMOTION1
	// AVCodecIdVmdvideo wraps AV_CODEC_ID_VMDVIDEO.
	AVCodecIdVmdvideo AVCodecID = C.AV_CODEC_ID_VMDVIDEO
	// AVCodecIdMszh wraps AV_CODEC_ID_MSZH.
	AVCodecIdMszh AVCodecID = C.AV_CODEC_ID_MSZH
	// AVCodecIdZlib wraps AV_CODEC_ID_ZLIB.
	AVCodecIdZlib AVCodecID = C.AV_CODEC_ID_ZLIB
	// AVCodecIdQtrle wraps AV_CODEC_ID_QTRLE.
	AVCodecIdQtrle AVCodecID = C.AV_CODEC_ID_QTRLE
	// AVCodecIdTscc wraps AV_CODEC_ID_TSCC.
	AVCodecIdTscc AVCodecID = C.AV_CODEC_ID_TSCC
	// AVCodecIdUlti wraps AV_CODEC_ID_ULTI.
	AVCodecIdUlti AVCodecID = C.AV_CODEC_ID_ULTI
	// AVCodecIdQdraw wraps AV_CODEC_ID_QDRAW.
	AVCodecIdQdraw AVCodecID = C.AV_CODEC_ID_QDRAW
	// AVCodecIdVixl wraps AV_CODEC_ID_VIXL.
	AVCodecIdVixl AVCodecID = C.AV_CODEC_ID_VIXL
	// AVCodecIdQpeg wraps AV_CODEC_ID_QPEG.
	AVCodecIdQpeg AVCodecID = C.AV_CODEC_ID_QPEG
	// AVCodecIdPng wraps AV_CODEC_ID_PNG.
	AVCodecIdPng AVCodecID = C.AV_CODEC_ID_PNG
	// AVCodecIdPpm wraps AV_CODEC_ID_PPM.
	AVCodecIdPpm AVCodecID = C.AV_CODEC_ID_PPM
	// AVCodecIdPbm wraps AV_CODEC_ID_PBM.
	AVCodecIdPbm AVCodecID = C.AV_CODEC_ID_PBM
	// AVCodecIdPgm wraps AV_CODEC_ID_PGM.
	AVCodecIdPgm AVCodecID = C.AV_CODEC_ID_PGM
	// AVCodecIdPgmyuv wraps AV_CODEC_ID_PGMYUV.
	AVCodecIdPgmyuv AVCodecID = C.AV_CODEC_ID_PGMYUV
	// AVCodecIdPam wraps AV_CODEC_ID_PAM.
	AVCodecIdPam AVCodecID = C.AV_CODEC_ID_PAM
	// AVCodecIdFfvhuff wraps AV_CODEC_ID_FFVHUFF.
	AVCodecIdFfvhuff AVCodecID = C.AV_CODEC_ID_FFVHUFF
	// AVCodecIdRv30 wraps AV_CODEC_ID_RV30.
	AVCodecIdRv30 AVCodecID = C.AV_CODEC_ID_RV30
	// AVCodecIdRv40 wraps AV_CODEC_ID_RV40.
	AVCodecIdRv40 AVCodecID = C.AV_CODEC_ID_RV40
	// AVCodecIdVc1 wraps AV_CODEC_ID_VC1.
	AVCodecIdVc1 AVCodecID = C.AV_CODEC_ID_VC1
	// AVCodecIdWmv3 wraps AV_CODEC_ID_WMV3.
	AVCodecIdWmv3 AVCodecID = C.AV_CODEC_ID_WMV3
	// AVCodecIdLoco wraps AV_CODEC_ID_LOCO.
	AVCodecIdLoco AVCodecID = C.AV_CODEC_ID_LOCO
	// AVCodecIdWnv1 wraps AV_CODEC_ID_WNV1.
	AVCodecIdWnv1 AVCodecID = C.AV_CODEC_ID_WNV1
	// AVCodecIdAasc wraps AV_CODEC_ID_AASC.
	AVCodecIdAasc AVCodecID = C.AV_CODEC_ID_AASC
	// AVCodecIdIndeo2 wraps AV_CODEC_ID_INDEO2.
	AVCodecIdIndeo2 AVCodecID = C.AV_CODEC_ID_INDEO2
	// AVCodecIdFraps wraps AV_CODEC_ID_FRAPS.
	AVCodecIdFraps AVCodecID = C.AV_CODEC_ID_FRAPS
	// AVCodecIdTruemotion2 wraps AV_CODEC_ID_TRUEMOTION2.
	AVCodecIdTruemotion2 AVCodecID = C.AV_CODEC_ID_TRUEMOTION2
	// AVCodecIdBmp wraps AV_CODEC_ID_BMP.
	AVCodecIdBmp AVCodecID = C.AV_CODEC_ID_BMP
	// AVCodecIdCscd wraps AV_CODEC_ID_CSCD.
	AVCodecIdCscd AVCodecID = C.AV_CODEC_ID_CSCD
	// AVCodecIdMmvideo wraps AV_CODEC_ID_MMVIDEO.
	AVCodecIdMmvideo AVCodecID = C.AV_CODEC_ID_MMVIDEO
	// AVCodecIdZmbv wraps AV_CODEC_ID_ZMBV.
	AVCodecIdZmbv AVCodecID = C.AV_CODEC_ID_ZMBV
	// AVCodecIdAvs wraps AV_CODEC_ID_AVS.
	AVCodecIdAvs AVCodecID = C.AV_CODEC_ID_AVS
	// AVCodecIdSmackvideo wraps AV_CODEC_ID_SMACKVIDEO.
	AVCodecIdSmackvideo AVCodecID = C.AV_CODEC_ID_SMACKVIDEO
	// AVCodecIdNuv wraps AV_CODEC_ID_NUV.
	AVCodecIdNuv AVCodecID = C.AV_CODEC_ID_NUV
	// AVCodecIdKmvc wraps AV_CODEC_ID_KMVC.
	AVCodecIdKmvc AVCodecID = C.AV_CODEC_ID_KMVC
	// AVCodecIdFlashsv wraps AV_CODEC_ID_FLASHSV.
	AVCodecIdFlashsv AVCodecID = C.AV_CODEC_ID_FLASHSV
	// AVCodecIdCavs wraps AV_CODEC_ID_CAVS.
	AVCodecIdCavs AVCodecID = C.AV_CODEC_ID_CAVS
	// AVCodecIdJpeg2000 wraps AV_CODEC_ID_JPEG2000.
	AVCodecIdJpeg2000 AVCodecID = C.AV_CODEC_ID_JPEG2000
	// AVCodecIdVmnc wraps AV_CODEC_ID_VMNC.
	AVCodecIdVmnc AVCodecID = C.AV_CODEC_ID_VMNC
	// AVCodecIdVp5 wraps AV_CODEC_ID_VP5.
	AVCodecIdVp5 AVCodecID = C.AV_CODEC_ID_VP5
	// AVCodecIdVp6 wraps AV_CODEC_ID_VP6.
	AVCodecIdVp6 AVCodecID = C.AV_CODEC_ID_VP6
	// AVCodecIdVp6F wraps AV_CODEC_ID_VP6F.
	AVCodecIdVp6F AVCodecID = C.AV_CODEC_ID_VP6F
	// AVCodecIdTarga wraps AV_CODEC_ID_TARGA.
	AVCodecIdTarga AVCodecID = C.AV_CODEC_ID_TARGA
	// AVCodecIdDsicinvideo wraps AV_CODEC_ID_DSICINVIDEO.
	AVCodecIdDsicinvideo AVCodecID = C.AV_CODEC_ID_DSICINVIDEO
	// AVCodecIdTiertexseqvideo wraps AV_CODEC_ID_TIERTEXSEQVIDEO.
	AVCodecIdTiertexseqvideo AVCodecID = C.AV_CODEC_ID_TIERTEXSEQVIDEO
	// AVCodecIdTiff wraps AV_CODEC_ID_TIFF.
	AVCodecIdTiff AVCodecID = C.AV_CODEC_ID_TIFF
	// AVCodecIdGif wraps AV_CODEC_ID_GIF.
	AVCodecIdGif AVCodecID = C.AV_CODEC_ID_GIF
	// AVCodecIdDxa wraps AV_CODEC_ID_DXA.
	AVCodecIdDxa AVCodecID = C.AV_CODEC_ID_DXA
	// AVCodecIdDnxhd wraps AV_CODEC_ID_DNXHD.
	AVCodecIdDnxhd AVCodecID = C.AV_CODEC_ID_DNXHD
	// AVCodecIdThp wraps AV_CODEC_ID_THP.
	AVCodecIdThp AVCodecID = C.AV_CODEC_ID_THP
	// AVCodecIdSgi wraps AV_CODEC_ID_SGI.
	AVCodecIdSgi AVCodecID = C.AV_CODEC_ID_SGI
	// AVCodecIdC93 wraps AV_CODEC_ID_C93.
	AVCodecIdC93 AVCodecID = C.AV_CODEC_ID_C93
	// AVCodecIdBethsoftvid wraps AV_CODEC_ID_BETHSOFTVID.
	AVCodecIdBethsoftvid AVCodecID = C.AV_CODEC_ID_BETHSOFTVID
	// AVCodecIdPtx wraps AV_CODEC_ID_PTX.
	AVCodecIdPtx AVCodecID = C.AV_CODEC_ID_PTX
	// AVCodecIdTxd wraps AV_CODEC_ID_TXD.
	AVCodecIdTxd AVCodecID = C.AV_CODEC_ID_TXD
	// AVCodecIdVp6A wraps AV_CODEC_ID_VP6A.
	AVCodecIdVp6A AVCodecID = C.AV_CODEC_ID_VP6A
	// AVCodecIdAmv wraps AV_CODEC_ID_AMV.
	AVCodecIdAmv AVCodecID = C.AV_CODEC_ID_AMV
	// AVCodecIdVb wraps AV_CODEC_ID_VB.
	AVCodecIdVb AVCodecID = C.AV_CODEC_ID_VB
	// AVCodecIdPcx wraps AV_CODEC_ID_PCX.
	AVCodecIdPcx AVCodecID = C.AV_CODEC_ID_PCX
	// AVCodecIdSunrast wraps AV_CODEC_ID_SUNRAST.
	AVCodecIdSunrast AVCodecID = C.AV_CODEC_ID_SUNRAST
	// AVCodecIdIndeo4 wraps AV_CODEC_ID_INDEO4.
	AVCodecIdIndeo4 AVCodecID = C.AV_CODEC_ID_INDEO4
	// AVCodecIdIndeo5 wraps AV_CODEC_ID_INDEO5.
	AVCodecIdIndeo5 AVCodecID = C.AV_CODEC_ID_INDEO5
	// AVCodecIdMimic wraps AV_CODEC_ID_MIMIC.
	AVCodecIdMimic AVCodecID = C.AV_CODEC_ID_MIMIC
	// AVCodecIdRl2 wraps AV_CODEC_ID_RL2.
	AVCodecIdRl2 AVCodecID = C.AV_CODEC_ID_RL2
	// AVCodecIdEscape124 wraps AV_CODEC_ID_ESCAPE124.
	AVCodecIdEscape124 AVCodecID = C.AV_CODEC_ID_ESCAPE124
	// AVCodecIdDirac wraps AV_CODEC_ID_DIRAC.
	AVCodecIdDirac AVCodecID = C.AV_CODEC_ID_DIRAC
	// AVCodecIdBfi wraps AV_CODEC_ID_BFI.
	AVCodecIdBfi AVCodecID = C.AV_CODEC_ID_BFI
	// AVCodecIdCmv wraps AV_CODEC_ID_CMV.
	AVCodecIdCmv AVCodecID = C.AV_CODEC_ID_CMV
	// AVCodecIdMotionpixels wraps AV_CODEC_ID_MOTIONPIXELS.
	AVCodecIdMotionpixels AVCodecID = C.AV_CODEC_ID_MOTIONPIXELS
	// AVCodecIdTgv wraps AV_CODEC_ID_TGV.
	AVCodecIdTgv AVCodecID = C.AV_CODEC_ID_TGV
	// AVCodecIdTgq wraps AV_CODEC_ID_TGQ.
	AVCodecIdTgq AVCodecID = C.AV_CODEC_ID_TGQ
	// AVCodecIdTqi wraps AV_CODEC_ID_TQI.
	AVCodecIdTqi AVCodecID = C.AV_CODEC_ID_TQI
	// AVCodecIdAura wraps AV_CODEC_ID_AURA.
	AVCodecIdAura AVCodecID = C.AV_CODEC_ID_AURA
	// AVCodecIdAura2 wraps AV_CODEC_ID_AURA2.
	AVCodecIdAura2 AVCodecID = C.AV_CODEC_ID_AURA2
	// AVCodecIdV210X wraps AV_CODEC_ID_V210X.
	AVCodecIdV210X AVCodecID = C.AV_CODEC_ID_V210X
	// AVCodecIdTmv wraps AV_CODEC_ID_TMV.
	AVCodecIdTmv AVCodecID = C.AV_CODEC_ID_TMV
	// AVCodecIdV210 wraps AV_CODEC_ID_V210.
	AVCodecIdV210 AVCodecID = C.AV_CODEC_ID_V210
	// AVCodecIdDpx wraps AV_CODEC_ID_DPX.
	AVCodecIdDpx AVCodecID = C.AV_CODEC_ID_DPX
	// AVCodecIdMad wraps AV_CODEC_ID_MAD.
	AVCodecIdMad AVCodecID = C.AV_CODEC_ID_MAD
	// AVCodecIdFrwu wraps AV_CODEC_ID_FRWU.
	AVCodecIdFrwu AVCodecID = C.AV_CODEC_ID_FRWU
	// AVCodecIdFlashsv2 wraps AV_CODEC_ID_FLASHSV2.
	AVCodecIdFlashsv2 AVCodecID = C.AV_CODEC_ID_FLASHSV2
	// AVCodecIdCdgraphics wraps AV_CODEC_ID_CDGRAPHICS.
	AVCodecIdCdgraphics AVCodecID = C.AV_CODEC_ID_CDGRAPHICS
	// AVCodecIdR210 wraps AV_CODEC_ID_R210.
	AVCodecIdR210 AVCodecID = C.AV_CODEC_ID_R210
	// AVCodecIdAnm wraps AV_CODEC_ID_ANM.
	AVCodecIdAnm AVCodecID = C.AV_CODEC_ID_ANM
	// AVCodecIdBinkvideo wraps AV_CODEC_ID_BINKVIDEO.
	AVCodecIdBinkvideo AVCodecID = C.AV_CODEC_ID_BINKVIDEO
	// AVCodecIdIffIlbm wraps AV_CODEC_ID_IFF_ILBM.
	AVCodecIdIffIlbm AVCodecID = C.AV_CODEC_ID_IFF_ILBM
	// AVCodecIdKgv1 wraps AV_CODEC_ID_KGV1.
	AVCodecIdKgv1 AVCodecID = C.AV_CODEC_ID_KGV1
	// AVCodecIdYop wraps AV_CODEC_ID_YOP.
	AVCodecIdYop AVCodecID = C.AV_CODEC_ID_YOP
	// AVCodecIdVp8 wraps AV_CODEC_ID_VP8.
	AVCodecIdVp8 AVCodecID = C.AV_CODEC_ID_VP8
	// AVCodecIdPictor wraps AV_CODEC_ID_PICTOR.
	AVCodecIdPictor AVCodecID = C.AV_CODEC_ID_PICTOR
	// AVCodecIdAnsi wraps AV_CODEC_ID_ANSI.
	AVCodecIdAnsi AVCodecID = C.AV_CODEC_ID_ANSI
	// AVCodecIdA64Multi wraps AV_CODEC_ID_A64_MULTI.
	AVCodecIdA64Multi AVCodecID = C.AV_CODEC_ID_A64_MULTI
	// AVCodecIdA64Multi5 wraps AV_CODEC_ID_A64_MULTI5.
	AVCodecIdA64Multi5 AVCodecID = C.AV_CODEC_ID_A64_MULTI5
	// AVCodecIdR10K wraps AV_CODEC_ID_R10K.
	AVCodecIdR10K AVCodecID = C.AV_CODEC_ID_R10K
	// AVCodecIdMxpeg wraps AV_CODEC_ID_MXPEG.
	AVCodecIdMxpeg AVCodecID = C.AV_CODEC_ID_MXPEG
	// AVCodecIdLagarith wraps AV_CODEC_ID_LAGARITH.
	AVCodecIdLagarith AVCodecID = C.AV_CODEC_ID_LAGARITH
	// AVCodecIdProres wraps AV_CODEC_ID_PRORES.
	AVCodecIdProres AVCodecID = C.AV_CODEC_ID_PRORES
	// AVCodecIdJv wraps AV_CODEC_ID_JV.
	AVCodecIdJv AVCodecID = C.AV_CODEC_ID_JV
	// AVCodecIdDfa wraps AV_CODEC_ID_DFA.
	AVCodecIdDfa AVCodecID = C.AV_CODEC_ID_DFA
	// AVCodecIdWmv3Image wraps AV_CODEC_ID_WMV3IMAGE.
	AVCodecIdWmv3Image AVCodecID = C.AV_CODEC_ID_WMV3IMAGE
	// AVCodecIdVc1Image wraps AV_CODEC_ID_VC1IMAGE.
	AVCodecIdVc1Image AVCodecID = C.AV_CODEC_ID_VC1IMAGE
	// AVCodecIdUtvideo wraps AV_CODEC_ID_UTVIDEO.
	AVCodecIdUtvideo AVCodecID = C.AV_CODEC_ID_UTVIDEO
	// AVCodecIdBmvVideo wraps AV_CODEC_ID_BMV_VIDEO.
	AVCodecIdBmvVideo AVCodecID = C.AV_CODEC_ID_BMV_VIDEO
	// AVCodecIdVble wraps AV_CODEC_ID_VBLE.
	AVCodecIdVble AVCodecID = C.AV_CODEC_ID_VBLE
	// AVCodecIdDxtory wraps AV_CODEC_ID_DXTORY.
	AVCodecIdDxtory AVCodecID = C.AV_CODEC_ID_DXTORY
	// AVCodecIdV410 wraps AV_CODEC_ID_V410.
	AVCodecIdV410 AVCodecID = C.AV_CODEC_ID_V410
	// AVCodecIdXwd wraps AV_CODEC_ID_XWD.
	AVCodecIdXwd AVCodecID = C.AV_CODEC_ID_XWD
	// AVCodecIdCdxl wraps AV_CODEC_ID_CDXL.
	AVCodecIdCdxl AVCodecID = C.AV_CODEC_ID_CDXL
	// AVCodecIdXbm wraps AV_CODEC_ID_XBM.
	AVCodecIdXbm AVCodecID = C.AV_CODEC_ID_XBM
	// AVCodecIdZerocodec wraps AV_CODEC_ID_ZEROCODEC.
	AVCodecIdZerocodec AVCodecID = C.AV_CODEC_ID_ZEROCODEC
	// AVCodecIdMss1 wraps AV_CODEC_ID_MSS1.
	AVCodecIdMss1 AVCodecID = C.AV_CODEC_ID_MSS1
	// AVCodecIdMsa1 wraps AV_CODEC_ID_MSA1.
	AVCodecIdMsa1 AVCodecID = C.AV_CODEC_ID_MSA1
	// AVCodecIdTscc2 wraps AV_CODEC_ID_TSCC2.
	AVCodecIdTscc2 AVCodecID = C.AV_CODEC_ID_TSCC2
	// AVCodecIdMts2 wraps AV_CODEC_ID_MTS2.
	AVCodecIdMts2 AVCodecID = C.AV_CODEC_ID_MTS2
	// AVCodecIdCllc wraps AV_CODEC_ID_CLLC.
	AVCodecIdCllc AVCodecID = C.AV_CODEC_ID_CLLC
	// AVCodecIdMss2 wraps AV_CODEC_ID_MSS2.
	AVCodecIdMss2 AVCodecID = C.AV_CODEC_ID_MSS2
	// AVCodecIdVp9 wraps AV_CODEC_ID_VP9.
	AVCodecIdVp9 AVCodecID = C.AV_CODEC_ID_VP9
	// AVCodecIdAic wraps AV_CODEC_ID_AIC.
	AVCodecIdAic AVCodecID = C.AV_CODEC_ID_AIC
	// AVCodecIdEscape130 wraps AV_CODEC_ID_ESCAPE130.
	AVCodecIdEscape130 AVCodecID = C.AV_CODEC_ID_ESCAPE130
	// AVCodecIdG2M wraps AV_CODEC_ID_G2M.
	AVCodecIdG2M AVCodecID = C.AV_CODEC_ID_G2M
	// AVCodecIdWebp wraps AV_CODEC_ID_WEBP.
	AVCodecIdWebp AVCodecID = C.AV_CODEC_ID_WEBP
	// AVCodecIdHnm4Video wraps AV_CODEC_ID_HNM4_VIDEO.
	AVCodecIdHnm4Video AVCodecID = C.AV_CODEC_ID_HNM4_VIDEO
	// AVCodecIdHevc wraps AV_CODEC_ID_HEVC.
	AVCodecIdHevc AVCodecID = C.AV_CODEC_ID_HEVC
	// AVCodecIdFic wraps AV_CODEC_ID_FIC.
	AVCodecIdFic AVCodecID = C.AV_CODEC_ID_FIC
	// AVCodecIdAliasPix wraps AV_CODEC_ID_ALIAS_PIX.
	AVCodecIdAliasPix AVCodecID = C.AV_CODEC_ID_ALIAS_PIX
	// AVCodecIdBrenderPix wraps AV_CODEC_ID_BRENDER_PIX.
	AVCodecIdBrenderPix AVCodecID = C.AV_CODEC_ID_BRENDER_PIX
	// AVCodecIdPafVideo wraps AV_CODEC_ID_PAF_VIDEO.
	AVCodecIdPafVideo AVCodecID = C.AV_CODEC_ID_PAF_VIDEO
	// AVCodecIdExr wraps AV_CODEC_ID_EXR.
	AVCodecIdExr AVCodecID = C.AV_CODEC_ID_EXR
	// AVCodecIdVp7 wraps AV_CODEC_ID_VP7.
	AVCodecIdVp7 AVCodecID = C.AV_CODEC_ID_VP7
	// AVCodecIdSanm wraps AV_CODEC_ID_SANM.
	AVCodecIdSanm AVCodecID = C.AV_CODEC_ID_SANM
	// AVCodecIdSgirle wraps AV_CODEC_ID_SGIRLE.
	AVCodecIdSgirle AVCodecID = C.AV_CODEC_ID_SGIRLE
	// AVCodecIdMvc1 wraps AV_CODEC_ID_MVC1.
	AVCodecIdMvc1 AVCodecID = C.AV_CODEC_ID_MVC1
	// AVCodecIdMvc2 wraps AV_CODEC_ID_MVC2.
	AVCodecIdMvc2 AVCodecID = C.AV_CODEC_ID_MVC2
	// AVCodecIdHqx wraps AV_CODEC_ID_HQX.
	AVCodecIdHqx AVCodecID = C.AV_CODEC_ID_HQX
	// AVCodecIdTdsc wraps AV_CODEC_ID_TDSC.
	AVCodecIdTdsc AVCodecID = C.AV_CODEC_ID_TDSC
	// AVCodecIdHqHqa wraps AV_CODEC_ID_HQ_HQA.
	AVCodecIdHqHqa AVCodecID = C.AV_CODEC_ID_HQ_HQA
	// AVCodecIdHap wraps AV_CODEC_ID_HAP.
	AVCodecIdHap AVCodecID = C.AV_CODEC_ID_HAP
	// AVCodecIdDds wraps AV_CODEC_ID_DDS.
	AVCodecIdDds AVCodecID = C.AV_CODEC_ID_DDS
	// AVCodecIdDxv wraps AV_CODEC_ID_DXV.
	AVCodecIdDxv AVCodecID = C.AV_CODEC_ID_DXV
	// AVCodecIdScreenpresso wraps AV_CODEC_ID_SCREENPRESSO.
	AVCodecIdScreenpresso AVCodecID = C.AV_CODEC_ID_SCREENPRESSO
	// AVCodecIdRscc wraps AV_CODEC_ID_RSCC.
	AVCodecIdRscc AVCodecID = C.AV_CODEC_ID_RSCC
	// AVCodecIdAvs2 wraps AV_CODEC_ID_AVS2.
	AVCodecIdAvs2 AVCodecID = C.AV_CODEC_ID_AVS2
	// AVCodecIdPgx wraps AV_CODEC_ID_PGX.
	AVCodecIdPgx AVCodecID = C.AV_CODEC_ID_PGX
	// AVCodecIdAvs3 wraps AV_CODEC_ID_AVS3.
	AVCodecIdAvs3 AVCodecID = C.AV_CODEC_ID_AVS3
	// AVCodecIdMsp2 wraps AV_CODEC_ID_MSP2.
	AVCodecIdMsp2 AVCodecID = C.AV_CODEC_ID_MSP2
	// AVCodecIdVvc wraps AV_CODEC_ID_VVC.
	AVCodecIdVvc AVCodecID = C.AV_CODEC_ID_VVC
	// AVCodecIdY41P wraps AV_CODEC_ID_Y41P.
	AVCodecIdY41P AVCodecID = C.AV_CODEC_ID_Y41P
	// AVCodecIdAvrp wraps AV_CODEC_ID_AVRP.
	AVCodecIdAvrp AVCodecID = C.AV_CODEC_ID_AVRP
	// AVCodecId012V wraps AV_CODEC_ID_012V.
	AVCodecId012V AVCodecID = C.AV_CODEC_ID_012V
	// AVCodecIdAvui wraps AV_CODEC_ID_AVUI.
	AVCodecIdAvui AVCodecID = C.AV_CODEC_ID_AVUI
	// AVCodecIdAyuv wraps AV_CODEC_ID_AYUV.
	AVCodecIdAyuv AVCodecID = C.AV_CODEC_ID_AYUV
	// AVCodecIdTargaY216 wraps AV_CODEC_ID_TARGA_Y216.
	AVCodecIdTargaY216 AVCodecID = C.AV_CODEC_ID_TARGA_Y216
	// AVCodecIdV308 wraps AV_CODEC_ID_V308.
	AVCodecIdV308 AVCodecID = C.AV_CODEC_ID_V308
	// AVCodecIdV408 wraps AV_CODEC_ID_V408.
	AVCodecIdV408 AVCodecID = C.AV_CODEC_ID_V408
	// AVCodecIdYuv4 wraps AV_CODEC_ID_YUV4.
	AVCodecIdYuv4 AVCodecID = C.AV_CODEC_ID_YUV4
	// AVCodecIdAvrn wraps AV_CODEC_ID_AVRN.
	AVCodecIdAvrn AVCodecID = C.AV_CODEC_ID_AVRN
	// AVCodecIdCpia wraps AV_CODEC_ID_CPIA.
	AVCodecIdCpia AVCodecID = C.AV_CODEC_ID_CPIA
	// AVCodecIdXface wraps AV_CODEC_ID_XFACE.
	AVCodecIdXface AVCodecID = C.AV_CODEC_ID_XFACE
	// AVCodecIdSnow wraps AV_CODEC_ID_SNOW.
	AVCodecIdSnow AVCodecID = C.AV_CODEC_ID_SNOW
	// AVCodecIdSmvjpeg wraps AV_CODEC_ID_SMVJPEG.
	AVCodecIdSmvjpeg AVCodecID = C.AV_CODEC_ID_SMVJPEG
	// AVCodecIdApng wraps AV_CODEC_ID_APNG.
	AVCodecIdApng AVCodecID = C.AV_CODEC_ID_APNG
	// AVCodecIdDaala wraps AV_CODEC_ID_DAALA.
	AVCodecIdDaala AVCodecID = C.AV_CODEC_ID_DAALA
	// AVCodecIdCfhd wraps AV_CODEC_ID_CFHD.
	AVCodecIdCfhd AVCodecID = C.AV_CODEC_ID_CFHD
	// AVCodecIdTruemotion2Rt wraps AV_CODEC_ID_TRUEMOTION2RT.
	AVCodecIdTruemotion2Rt AVCodecID = C.AV_CODEC_ID_TRUEMOTION2RT
	// AVCodecIdM101 wraps AV_CODEC_ID_M101.
	AVCodecIdM101 AVCodecID = C.AV_CODEC_ID_M101
	// AVCodecIdMagicyuv wraps AV_CODEC_ID_MAGICYUV.
	AVCodecIdMagicyuv AVCodecID = C.AV_CODEC_ID_MAGICYUV
	// AVCodecIdSheervideo wraps AV_CODEC_ID_SHEERVIDEO.
	AVCodecIdSheervideo AVCodecID = C.AV_CODEC_ID_SHEERVIDEO
	// AVCodecIdYlc wraps AV_CODEC_ID_YLC.
	AVCodecIdYlc AVCodecID = C.AV_CODEC_ID_YLC
	// AVCodecIdPsd wraps AV_CODEC_ID_PSD.
	AVCodecIdPsd AVCodecID = C.AV_CODEC_ID_PSD
	// AVCodecIdPixlet wraps AV_CODEC_ID_PIXLET.
	AVCodecIdPixlet AVCodecID = C.AV_CODEC_ID_PIXLET
	// AVCodecIdSpeedhq wraps AV_CODEC_ID_SPEEDHQ.
	AVCodecIdSpeedhq AVCodecID = C.AV_CODEC_ID_SPEEDHQ
	// AVCodecIdFmvc wraps AV_CODEC_ID_FMVC.
	AVCodecIdFmvc AVCodecID = C.AV_CODEC_ID_FMVC
	// AVCodecIdScpr wraps AV_CODEC_ID_SCPR.
	AVCodecIdScpr AVCodecID = C.AV_CODEC_ID_SCPR
	// AVCodecIdClearvideo wraps AV_CODEC_ID_CLEARVIDEO.
	AVCodecIdClearvideo AVCodecID = C.AV_CODEC_ID_CLEARVIDEO
	// AVCodecIdXpm wraps AV_CODEC_ID_XPM.
	AVCodecIdXpm AVCodecID = C.AV_CODEC_ID_XPM
	// AVCodecIdAv1 wraps AV_CODEC_ID_AV1.
	AVCodecIdAv1 AVCodecID = C.AV_CODEC_ID_AV1
	// AVCodecIdBitpacked wraps AV_CODEC_ID_BITPACKED.
	AVCodecIdBitpacked AVCodecID = C.AV_CODEC_ID_BITPACKED
	// AVCodecIdMscc wraps AV_CODEC_ID_MSCC.
	AVCodecIdMscc AVCodecID = C.AV_CODEC_ID_MSCC
	// AVCodecIdSrgc wraps AV_CODEC_ID_SRGC.
	AVCodecIdSrgc AVCodecID = C.AV_CODEC_ID_SRGC
	// AVCodecIdSvg wraps AV_CODEC_ID_SVG.
	AVCodecIdSvg AVCodecID = C.AV_CODEC_ID_SVG
	// AVCodecIdGdv wraps AV_CODEC_ID_GDV.
	AVCodecIdGdv AVCodecID = C.AV_CODEC_ID_GDV
	// AVCodecIdFits wraps AV_CODEC_ID_FITS.
	AVCodecIdFits AVCodecID = C.AV_CODEC_ID_FITS
	// AVCodecIdImm4 wraps AV_CODEC_ID_IMM4.
	AVCodecIdImm4 AVCodecID = C.AV_CODEC_ID_IMM4
	// AVCodecIdProsumer wraps AV_CODEC_ID_PROSUMER.
	AVCodecIdProsumer AVCodecID = C.AV_CODEC_ID_PROSUMER
	// AVCodecIdMwsc wraps AV_CODEC_ID_MWSC.
	AVCodecIdMwsc AVCodecID = C.AV_CODEC_ID_MWSC
	// AVCodecIdWcmv wraps AV_CODEC_ID_WCMV.
	AVCodecIdWcmv AVCodecID = C.AV_CODEC_ID_WCMV
	// AVCodecIdRasc wraps AV_CODEC_ID_RASC.
	AVCodecIdRasc AVCodecID = C.AV_CODEC_ID_RASC
	// AVCodecIdHymt wraps AV_CODEC_ID_HYMT.
	AVCodecIdHymt AVCodecID = C.AV_CODEC_ID_HYMT
	// AVCodecIdArbc wraps AV_CODEC_ID_ARBC.
	AVCodecIdArbc AVCodecID = C.AV_CODEC_ID_ARBC
	// AVCodecIdAgm wraps AV_CODEC_ID_AGM.
	AVCodecIdAgm AVCodecID = C.AV_CODEC_ID_AGM
	// AVCodecIdLscr wraps AV_CODEC_ID_LSCR.
	AVCodecIdLscr AVCodecID = C.AV_CODEC_ID_LSCR
	// AVCodecIdVp4 wraps AV_CODEC_ID_VP4.
	AVCodecIdVp4 AVCodecID = C.AV_CODEC_ID_VP4
	// AVCodecIdImm5 wraps AV_CODEC_ID_IMM5.
	AVCodecIdImm5 AVCodecID = C.AV_CODEC_ID_IMM5
	// AVCodecIdMvdv wraps AV_CODEC_ID_MVDV.
	AVCodecIdMvdv AVCodecID = C.AV_CODEC_ID_MVDV
	// AVCodecIdMvha wraps AV_CODEC_ID_MVHA.
	AVCodecIdMvha AVCodecID = C.AV_CODEC_ID_MVHA
	// AVCodecIdCdtoons wraps AV_CODEC_ID_CDTOONS.
	AVCodecIdCdtoons AVCodecID = C.AV_CODEC_ID_CDTOONS
	// AVCodecIdMv30 wraps AV_CODEC_ID_MV30.
	AVCodecIdMv30 AVCodecID = C.AV_CODEC_ID_MV30
	// AVCodecIdNotchlc wraps AV_CODEC_ID_NOTCHLC.
	AVCodecIdNotchlc AVCodecID = C.AV_CODEC_ID_NOTCHLC
	// AVCodecIdPfm wraps AV_CODEC_ID_PFM.
	AVCodecIdPfm AVCodecID = C.AV_CODEC_ID_PFM
	// AVCodecIdMobiclip wraps AV_CODEC_ID_MOBICLIP.
	AVCodecIdMobiclip AVCodecID = C.AV_CODEC_ID_MOBICLIP
	// AVCodecIdPhotocd wraps AV_CODEC_ID_PHOTOCD.
	AVCodecIdPhotocd AVCodecID = C.AV_CODEC_ID_PHOTOCD
	// AVCodecIdIpu wraps AV_CODEC_ID_IPU.
	AVCodecIdIpu AVCodecID = C.AV_CODEC_ID_IPU
	// AVCodecIdArgo wraps AV_CODEC_ID_ARGO.
	AVCodecIdArgo AVCodecID = C.AV_CODEC_ID_ARGO
	// AVCodecIdCri wraps AV_CODEC_ID_CRI.
	AVCodecIdCri AVCodecID = C.AV_CODEC_ID_CRI
	// AVCodecIdSimbiosisImx wraps AV_CODEC_ID_SIMBIOSIS_IMX.
	AVCodecIdSimbiosisImx AVCodecID = C.AV_CODEC_ID_SIMBIOSIS_IMX
	// AVCodecIdSgaVideo wraps AV_CODEC_ID_SGA_VIDEO.
	AVCodecIdSgaVideo AVCodecID = C.AV_CODEC_ID_SGA_VIDEO
	// AVCodecIdGem wraps AV_CODEC_ID_GEM.
	AVCodecIdGem AVCodecID = C.AV_CODEC_ID_GEM
	// AVCodecIdVbn wraps AV_CODEC_ID_VBN.
	AVCodecIdVbn AVCodecID = C.AV_CODEC_ID_VBN
	// AVCodecIdJpegxl wraps AV_CODEC_ID_JPEGXL.
	AVCodecIdJpegxl AVCodecID = C.AV_CODEC_ID_JPEGXL
	// AVCodecIdQoi wraps AV_CODEC_ID_QOI.
	AVCodecIdQoi AVCodecID = C.AV_CODEC_ID_QOI
	// AVCodecIdPhm wraps AV_CODEC_ID_PHM.
	AVCodecIdPhm AVCodecID = C.AV_CODEC_ID_PHM
	// AVCodecIdRadianceHdr wraps AV_CODEC_ID_RADIANCE_HDR.
	AVCodecIdRadianceHdr AVCodecID = C.AV_CODEC_ID_RADIANCE_HDR
	// AVCodecIdWbmp wraps AV_CODEC_ID_WBMP.
	AVCodecIdWbmp AVCodecID = C.AV_CODEC_ID_WBMP
	// AVCodecIdMedia100 wraps AV_CODEC_ID_MEDIA100.
	AVCodecIdMedia100 AVCodecID = C.AV_CODEC_ID_MEDIA100
	// AVCodecIdVqc wraps AV_CODEC_ID_VQC.
	AVCodecIdVqc AVCodecID = C.AV_CODEC_ID_VQC
	// AVCodecIdFirstAudio wraps AV_CODEC_ID_FIRST_AUDIO.
	AVCodecIdFirstAudio AVCodecID = C.AV_CODEC_ID_FIRST_AUDIO
	// AVCodecIdPcmS16Le wraps AV_CODEC_ID_PCM_S16LE.
	AVCodecIdPcmS16Le AVCodecID = C.AV_CODEC_ID_PCM_S16LE
	// AVCodecIdPcmS16Be wraps AV_CODEC_ID_PCM_S16BE.
	AVCodecIdPcmS16Be AVCodecID = C.AV_CODEC_ID_PCM_S16BE
	// AVCodecIdPcmU16Le wraps AV_CODEC_ID_PCM_U16LE.
	AVCodecIdPcmU16Le AVCodecID = C.AV_CODEC_ID_PCM_U16LE
	// AVCodecIdPcmU16Be wraps AV_CODEC_ID_PCM_U16BE.
	AVCodecIdPcmU16Be AVCodecID = C.AV_CODEC_ID_PCM_U16BE
	// AVCodecIdPcmS8 wraps AV_CODEC_ID_PCM_S8.
	AVCodecIdPcmS8 AVCodecID = C.AV_CODEC_ID_PCM_S8
	// AVCodecIdPcmU8 wraps AV_CODEC_ID_PCM_U8.
	AVCodecIdPcmU8 AVCodecID = C.AV_CODEC_ID_PCM_U8
	// AVCodecIdPcmMulaw wraps AV_CODEC_ID_PCM_MULAW.
	AVCodecIdPcmMulaw AVCodecID = C.AV_CODEC_ID_PCM_MULAW
	// AVCodecIdPcmAlaw wraps AV_CODEC_ID_PCM_ALAW.
	AVCodecIdPcmAlaw AVCodecID = C.AV_CODEC_ID_PCM_ALAW
	// AVCodecIdPcmS32Le wraps AV_CODEC_ID_PCM_S32LE.
	AVCodecIdPcmS32Le AVCodecID = C.AV_CODEC_ID_PCM_S32LE
	// AVCodecIdPcmS32Be wraps AV_CODEC_ID_PCM_S32BE.
	AVCodecIdPcmS32Be AVCodecID = C.AV_CODEC_ID_PCM_S32BE
	// AVCodecIdPcmU32Le wraps AV_CODEC_ID_PCM_U32LE.
	AVCodecIdPcmU32Le AVCodecID = C.AV_CODEC_ID_PCM_U32LE
	// AVCodecIdPcmU32Be wraps AV_CODEC_ID_PCM_U32BE.
	AVCodecIdPcmU32Be AVCodecID = C.AV_CODEC_ID_PCM_U32BE
	// AVCodecIdPcmS24Le wraps AV_CODEC_ID_PCM_S24LE.
	AVCodecIdPcmS24Le AVCodecID = C.AV_CODEC_ID_PCM_S24LE
	// AVCodecIdPcmS24Be wraps AV_CODEC_ID_PCM_S24BE.
	AVCodecIdPcmS24Be AVCodecID = C.AV_CODEC_ID_PCM_S24BE
	// AVCodecIdPcmU24Le wraps AV_CODEC_ID_PCM_U24LE.
	AVCodecIdPcmU24Le AVCodecID = C.AV_CODEC_ID_PCM_U24LE
	// AVCodecIdPcmU24Be wraps AV_CODEC_ID_PCM_U24BE.
	AVCodecIdPcmU24Be AVCodecID = C.AV_CODEC_ID_PCM_U24BE
	// AVCodecIdPcmS24Daud wraps AV_CODEC_ID_PCM_S24DAUD.
	AVCodecIdPcmS24Daud AVCodecID = C.AV_CODEC_ID_PCM_S24DAUD
	// AVCodecIdPcmZork wraps AV_CODEC_ID_PCM_ZORK.
	AVCodecIdPcmZork AVCodecID = C.AV_CODEC_ID_PCM_ZORK
	// AVCodecIdPcmS16LePlanar wraps AV_CODEC_ID_PCM_S16LE_PLANAR.
	AVCodecIdPcmS16LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S16LE_PLANAR
	// AVCodecIdPcmDvd wraps AV_CODEC_ID_PCM_DVD.
	AVCodecIdPcmDvd AVCodecID = C.AV_CODEC_ID_PCM_DVD
	// AVCodecIdPcmF32Be wraps AV_CODEC_ID_PCM_F32BE.
	AVCodecIdPcmF32Be AVCodecID = C.AV_CODEC_ID_PCM_F32BE
	// AVCodecIdPcmF32Le wraps AV_CODEC_ID_PCM_F32LE.
	AVCodecIdPcmF32Le AVCodecID = C.AV_CODEC_ID_PCM_F32LE
	// AVCodecIdPcmF64Be wraps AV_CODEC_ID_PCM_F64BE.
	AVCodecIdPcmF64Be AVCodecID = C.AV_CODEC_ID_PCM_F64BE
	// AVCodecIdPcmF64Le wraps AV_CODEC_ID_PCM_F64LE.
	AVCodecIdPcmF64Le AVCodecID = C.AV_CODEC_ID_PCM_F64LE
	// AVCodecIdPcmBluray wraps AV_CODEC_ID_PCM_BLURAY.
	AVCodecIdPcmBluray AVCodecID = C.AV_CODEC_ID_PCM_BLURAY
	// AVCodecIdPcmLxf wraps AV_CODEC_ID_PCM_LXF.
	AVCodecIdPcmLxf AVCodecID = C.AV_CODEC_ID_PCM_LXF
	// AVCodecIdS302M wraps AV_CODEC_ID_S302M.
	AVCodecIdS302M AVCodecID = C.AV_CODEC_ID_S302M
	// AVCodecIdPcmS8Planar wraps AV_CODEC_ID_PCM_S8_PLANAR.
	AVCodecIdPcmS8Planar AVCodecID = C.AV_CODEC_ID_PCM_S8_PLANAR
	// AVCodecIdPcmS24LePlanar wraps AV_CODEC_ID_PCM_S24LE_PLANAR.
	AVCodecIdPcmS24LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S24LE_PLANAR
	// AVCodecIdPcmS32LePlanar wraps AV_CODEC_ID_PCM_S32LE_PLANAR.
	AVCodecIdPcmS32LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S32LE_PLANAR
	// AVCodecIdPcmS16BePlanar wraps AV_CODEC_ID_PCM_S16BE_PLANAR.
	AVCodecIdPcmS16BePlanar AVCodecID = C.AV_CODEC_ID_PCM_S16BE_PLANAR
	// AVCodecIdPcmS64Le wraps AV_CODEC_ID_PCM_S64LE.
	AVCodecIdPcmS64Le AVCodecID = C.AV_CODEC_ID_PCM_S64LE
	// AVCodecIdPcmS64Be wraps AV_CODEC_ID_PCM_S64BE.
	AVCodecIdPcmS64Be AVCodecID = C.AV_CODEC_ID_PCM_S64BE
	// AVCodecIdPcmF16Le wraps AV_CODEC_ID_PCM_F16LE.
	AVCodecIdPcmF16Le AVCodecID = C.AV_CODEC_ID_PCM_F16LE
	// AVCodecIdPcmF24Le wraps AV_CODEC_ID_PCM_F24LE.
	AVCodecIdPcmF24Le AVCodecID = C.AV_CODEC_ID_PCM_F24LE
	// AVCodecIdPcmVidc wraps AV_CODEC_ID_PCM_VIDC.
	AVCodecIdPcmVidc AVCodecID = C.AV_CODEC_ID_PCM_VIDC
	// AVCodecIdPcmSga wraps AV_CODEC_ID_PCM_SGA.
	AVCodecIdPcmSga AVCodecID = C.AV_CODEC_ID_PCM_SGA
	// AVCodecIdAdpcmImaQt wraps AV_CODEC_ID_ADPCM_IMA_QT.
	AVCodecIdAdpcmImaQt AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_QT
	// AVCodecIdAdpcmImaWav wraps AV_CODEC_ID_ADPCM_IMA_WAV.
	AVCodecIdAdpcmImaWav AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_WAV
	// AVCodecIdAdpcmImaDk3 wraps AV_CODEC_ID_ADPCM_IMA_DK3.
	AVCodecIdAdpcmImaDk3 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DK3
	// AVCodecIdAdpcmImaDk4 wraps AV_CODEC_ID_ADPCM_IMA_DK4.
	AVCodecIdAdpcmImaDk4 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DK4
	// AVCodecIdAdpcmImaWs wraps AV_CODEC_ID_ADPCM_IMA_WS.
	AVCodecIdAdpcmImaWs AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_WS
	// AVCodecIdAdpcmImaSmjpeg wraps AV_CODEC_ID_ADPCM_IMA_SMJPEG.
	AVCodecIdAdpcmImaSmjpeg AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_SMJPEG
	// AVCodecIdAdpcmMs wraps AV_CODEC_ID_ADPCM_MS.
	AVCodecIdAdpcmMs AVCodecID = C.AV_CODEC_ID_ADPCM_MS
	// AVCodecIdAdpcm4Xm wraps AV_CODEC_ID_ADPCM_4XM.
	AVCodecIdAdpcm4Xm AVCodecID = C.AV_CODEC_ID_ADPCM_4XM
	// AVCodecIdAdpcmXa wraps AV_CODEC_ID_ADPCM_XA.
	AVCodecIdAdpcmXa AVCodecID = C.AV_CODEC_ID_ADPCM_XA
	// AVCodecIdAdpcmAdx wraps AV_CODEC_ID_ADPCM_ADX.
	AVCodecIdAdpcmAdx AVCodecID = C.AV_CODEC_ID_ADPCM_ADX
	// AVCodecIdAdpcmEa wraps AV_CODEC_ID_ADPCM_EA.
	AVCodecIdAdpcmEa AVCodecID = C.AV_CODEC_ID_ADPCM_EA
	// AVCodecIdAdpcmG726 wraps AV_CODEC_ID_ADPCM_G726.
	AVCodecIdAdpcmG726 AVCodecID = C.AV_CODEC_ID_ADPCM_G726
	// AVCodecIdAdpcmCt wraps AV_CODEC_ID_ADPCM_CT.
	AVCodecIdAdpcmCt AVCodecID = C.AV_CODEC_ID_ADPCM_CT
	// AVCodecIdAdpcmSwf wraps AV_CODEC_ID_ADPCM_SWF.
	AVCodecIdAdpcmSwf AVCodecID = C.AV_CODEC_ID_ADPCM_SWF
	// AVCodecIdAdpcmYamaha wraps AV_CODEC_ID_ADPCM_YAMAHA.
	AVCodecIdAdpcmYamaha AVCodecID = C.AV_CODEC_ID_ADPCM_YAMAHA
	// AVCodecIdAdpcmSbpro4 wraps AV_CODEC_ID_ADPCM_SBPRO_4.
	AVCodecIdAdpcmSbpro4 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_4
	// AVCodecIdAdpcmSbpro3 wraps AV_CODEC_ID_ADPCM_SBPRO_3.
	AVCodecIdAdpcmSbpro3 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_3
	// AVCodecIdAdpcmSbpro2 wraps AV_CODEC_ID_ADPCM_SBPRO_2.
	AVCodecIdAdpcmSbpro2 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_2
	// AVCodecIdAdpcmThp wraps AV_CODEC_ID_ADPCM_THP.
	AVCodecIdAdpcmThp AVCodecID = C.AV_CODEC_ID_ADPCM_THP
	// AVCodecIdAdpcmImaAmv wraps AV_CODEC_ID_ADPCM_IMA_AMV.
	AVCodecIdAdpcmImaAmv AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_AMV
	// AVCodecIdAdpcmEaR1 wraps AV_CODEC_ID_ADPCM_EA_R1.
	AVCodecIdAdpcmEaR1 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R1
	// AVCodecIdAdpcmEaR3 wraps AV_CODEC_ID_ADPCM_EA_R3.
	AVCodecIdAdpcmEaR3 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R3
	// AVCodecIdAdpcmEaR2 wraps AV_CODEC_ID_ADPCM_EA_R2.
	AVCodecIdAdpcmEaR2 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R2
	// AVCodecIdAdpcmImaEaSead wraps AV_CODEC_ID_ADPCM_IMA_EA_SEAD.
	AVCodecIdAdpcmImaEaSead AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD
	// AVCodecIdAdpcmImaEaEacs wraps AV_CODEC_ID_ADPCM_IMA_EA_EACS.
	AVCodecIdAdpcmImaEaEacs AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_EACS
	// AVCodecIdAdpcmEaXas wraps AV_CODEC_ID_ADPCM_EA_XAS.
	AVCodecIdAdpcmEaXas AVCodecID = C.AV_CODEC_ID_ADPCM_EA_XAS
	// AVCodecIdAdpcmEaMaxisXa wraps AV_CODEC_ID_ADPCM_EA_MAXIS_XA.
	AVCodecIdAdpcmEaMaxisXa AVCodecID = C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA
	// AVCodecIdAdpcmImaIss wraps AV_CODEC_ID_ADPCM_IMA_ISS.
	AVCodecIdAdpcmImaIss AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ISS
	// AVCodecIdAdpcmG722 wraps AV_CODEC_ID_ADPCM_G722.
	AVCodecIdAdpcmG722 AVCodecID = C.AV_CODEC_ID_ADPCM_G722
	// AVCodecIdAdpcmImaApc wraps AV_CODEC_ID_ADPCM_IMA_APC.
	AVCodecIdAdpcmImaApc AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_APC
	// AVCodecIdAdpcmVima wraps AV_CODEC_ID_ADPCM_VIMA.
	AVCodecIdAdpcmVima AVCodecID = C.AV_CODEC_ID_ADPCM_VIMA
	// AVCodecIdAdpcmAfc wraps AV_CODEC_ID_ADPCM_AFC.
	AVCodecIdAdpcmAfc AVCodecID = C.AV_CODEC_ID_ADPCM_AFC
	// AVCodecIdAdpcmImaOki wraps AV_CODEC_ID_ADPCM_IMA_OKI.
	AVCodecIdAdpcmImaOki AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_OKI
	// AVCodecIdAdpcmDtk wraps AV_CODEC_ID_ADPCM_DTK.
	AVCodecIdAdpcmDtk AVCodecID = C.AV_CODEC_ID_ADPCM_DTK
	// AVCodecIdAdpcmImaRad wraps AV_CODEC_ID_ADPCM_IMA_RAD.
	AVCodecIdAdpcmImaRad AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_RAD
	// AVCodecIdAdpcmG726Le wraps AV_CODEC_ID_ADPCM_G726LE.
	AVCodecIdAdpcmG726Le AVCodecID = C.AV_CODEC_ID_ADPCM_G726LE
	// AVCodecIdAdpcmThpLe wraps AV_CODEC_ID_ADPCM_THP_LE.
	AVCodecIdAdpcmThpLe AVCodecID = C.AV_CODEC_ID_ADPCM_THP_LE
	// AVCodecIdAdpcmPsx wraps AV_CODEC_ID_ADPCM_PSX.
	AVCodecIdAdpcmPsx AVCodecID = C.AV_CODEC_ID_ADPCM_PSX
	// AVCodecIdAdpcmAica wraps AV_CODEC_ID_ADPCM_AICA.
	AVCodecIdAdpcmAica AVCodecID = C.AV_CODEC_ID_ADPCM_AICA
	// AVCodecIdAdpcmImaDat4 wraps AV_CODEC_ID_ADPCM_IMA_DAT4.
	AVCodecIdAdpcmImaDat4 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DAT4
	// AVCodecIdAdpcmMtaf wraps AV_CODEC_ID_ADPCM_MTAF.
	AVCodecIdAdpcmMtaf AVCodecID = C.AV_CODEC_ID_ADPCM_MTAF
	// AVCodecIdAdpcmAgm wraps AV_CODEC_ID_ADPCM_AGM.
	AVCodecIdAdpcmAgm AVCodecID = C.AV_CODEC_ID_ADPCM_AGM
	// AVCodecIdAdpcmArgo wraps AV_CODEC_ID_ADPCM_ARGO.
	AVCodecIdAdpcmArgo AVCodecID = C.AV_CODEC_ID_ADPCM_ARGO
	// AVCodecIdAdpcmImaSsi wraps AV_CODEC_ID_ADPCM_IMA_SSI.
	AVCodecIdAdpcmImaSsi AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_SSI
	// AVCodecIdAdpcmZork wraps AV_CODEC_ID_ADPCM_ZORK.
	AVCodecIdAdpcmZork AVCodecID = C.AV_CODEC_ID_ADPCM_ZORK
	// AVCodecIdAdpcmImaApm wraps AV_CODEC_ID_ADPCM_IMA_APM.
	AVCodecIdAdpcmImaApm AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_APM
	// AVCodecIdAdpcmImaAlp wraps AV_CODEC_ID_ADPCM_IMA_ALP.
	AVCodecIdAdpcmImaAlp AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ALP
	// AVCodecIdAdpcmImaMtf wraps AV_CODEC_ID_ADPCM_IMA_MTF.
	AVCodecIdAdpcmImaMtf AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_MTF
	// AVCodecIdAdpcmImaCunning wraps AV_CODEC_ID_ADPCM_IMA_CUNNING.
	AVCodecIdAdpcmImaCunning AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_CUNNING
	// AVCodecIdAdpcmImaMoflex wraps AV_CODEC_ID_ADPCM_IMA_MOFLEX.
	AVCodecIdAdpcmImaMoflex AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_MOFLEX
	// AVCodecIdAdpcmImaAcorn wraps AV_CODEC_ID_ADPCM_IMA_ACORN.
	AVCodecIdAdpcmImaAcorn AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ACORN
	// AVCodecIdAdpcmXmd wraps AV_CODEC_ID_ADPCM_XMD.
	AVCodecIdAdpcmXmd AVCodecID = C.AV_CODEC_ID_ADPCM_XMD
	// AVCodecIdAmrNb wraps AV_CODEC_ID_AMR_NB.
	AVCodecIdAmrNb AVCodecID = C.AV_CODEC_ID_AMR_NB
	// AVCodecIdAmrWb wraps AV_CODEC_ID_AMR_WB.
	AVCodecIdAmrWb AVCodecID = C.AV_CODEC_ID_AMR_WB
	// AVCodecIdRa144 wraps AV_CODEC_ID_RA_144.
	AVCodecIdRa144 AVCodecID = C.AV_CODEC_ID_RA_144
	// AVCodecIdRa288 wraps AV_CODEC_ID_RA_288.
	AVCodecIdRa288 AVCodecID = C.AV_CODEC_ID_RA_288
	// AVCodecIdRoqDpcm wraps AV_CODEC_ID_ROQ_DPCM.
	AVCodecIdRoqDpcm AVCodecID = C.AV_CODEC_ID_ROQ_DPCM
	// AVCodecIdInterplayDpcm wraps AV_CODEC_ID_INTERPLAY_DPCM.
	AVCodecIdInterplayDpcm AVCodecID = C.AV_CODEC_ID_INTERPLAY_DPCM
	// AVCodecIdXanDpcm wraps AV_CODEC_ID_XAN_DPCM.
	AVCodecIdXanDpcm AVCodecID = C.AV_CODEC_ID_XAN_DPCM
	// AVCodecIdSolDpcm wraps AV_CODEC_ID_SOL_DPCM.
	AVCodecIdSolDpcm AVCodecID = C.AV_CODEC_ID_SOL_DPCM
	// AVCodecIdSdx2Dpcm wraps AV_CODEC_ID_SDX2_DPCM.
	AVCodecIdSdx2Dpcm AVCodecID = C.AV_CODEC_ID_SDX2_DPCM
	// AVCodecIdGremlinDpcm wraps AV_CODEC_ID_GREMLIN_DPCM.
	AVCodecIdGremlinDpcm AVCodecID = C.AV_CODEC_ID_GREMLIN_DPCM
	// AVCodecIdDerfDpcm wraps AV_CODEC_ID_DERF_DPCM.
	AVCodecIdDerfDpcm AVCodecID = C.AV_CODEC_ID_DERF_DPCM
	// AVCodecIdWadyDpcm wraps AV_CODEC_ID_WADY_DPCM.
	AVCodecIdWadyDpcm AVCodecID = C.AV_CODEC_ID_WADY_DPCM
	// AVCodecIdCbd2Dpcm wraps AV_CODEC_ID_CBD2_DPCM.
	AVCodecIdCbd2Dpcm AVCodecID = C.AV_CODEC_ID_CBD2_DPCM
	// AVCodecIdMp2 wraps AV_CODEC_ID_MP2.
	AVCodecIdMp2 AVCodecID = C.AV_CODEC_ID_MP2
	// AVCodecIdMp3 wraps AV_CODEC_ID_MP3.
	AVCodecIdMp3 AVCodecID = C.AV_CODEC_ID_MP3
	// AVCodecIdAac wraps AV_CODEC_ID_AAC.
	AVCodecIdAac AVCodecID = C.AV_CODEC_ID_AAC
	// AVCodecIdAc3 wraps AV_CODEC_ID_AC3.
	AVCodecIdAc3 AVCodecID = C.AV_CODEC_ID_AC3
	// AVCodecIdDts wraps AV_CODEC_ID_DTS.
	AVCodecIdDts AVCodecID = C.AV_CODEC_ID_DTS
	// AVCodecIdVorbis wraps AV_CODEC_ID_VORBIS.
	AVCodecIdVorbis AVCodecID = C.AV_CODEC_ID_VORBIS
	// AVCodecIdDvaudio wraps AV_CODEC_ID_DVAUDIO.
	AVCodecIdDvaudio AVCodecID = C.AV_CODEC_ID_DVAUDIO
	// AVCodecIdWmav1 wraps AV_CODEC_ID_WMAV1.
	AVCodecIdWmav1 AVCodecID = C.AV_CODEC_ID_WMAV1
	// AVCodecIdWmav2 wraps AV_CODEC_ID_WMAV2.
	AVCodecIdWmav2 AVCodecID = C.AV_CODEC_ID_WMAV2
	// AVCodecIdMace3 wraps AV_CODEC_ID_MACE3.
	AVCodecIdMace3 AVCodecID = C.AV_CODEC_ID_MACE3
	// AVCodecIdMace6 wraps AV_CODEC_ID_MACE6.
	AVCodecIdMace6 AVCodecID = C.AV_CODEC_ID_MACE6
	// AVCodecIdVmdaudio wraps AV_CODEC_ID_VMDAUDIO.
	AVCodecIdVmdaudio AVCodecID = C.AV_CODEC_ID_VMDAUDIO
	// AVCodecIdFlac wraps AV_CODEC_ID_FLAC.
	AVCodecIdFlac AVCodecID = C.AV_CODEC_ID_FLAC
	// AVCodecIdMp3Adu wraps AV_CODEC_ID_MP3ADU.
	AVCodecIdMp3Adu AVCodecID = C.AV_CODEC_ID_MP3ADU
	// AVCodecIdMp3On4 wraps AV_CODEC_ID_MP3ON4.
	AVCodecIdMp3On4 AVCodecID = C.AV_CODEC_ID_MP3ON4
	// AVCodecIdShorten wraps AV_CODEC_ID_SHORTEN.
	AVCodecIdShorten AVCodecID = C.AV_CODEC_ID_SHORTEN
	// AVCodecIdAlac wraps AV_CODEC_ID_ALAC.
	AVCodecIdAlac AVCodecID = C.AV_CODEC_ID_ALAC
	// AVCodecIdWestwoodSnd1 wraps AV_CODEC_ID_WESTWOOD_SND1.
	AVCodecIdWestwoodSnd1 AVCodecID = C.AV_CODEC_ID_WESTWOOD_SND1
	// AVCodecIdGsm wraps AV_CODEC_ID_GSM.
	AVCodecIdGsm AVCodecID = C.AV_CODEC_ID_GSM
	// AVCodecIdQdm2 wraps AV_CODEC_ID_QDM2.
	AVCodecIdQdm2 AVCodecID = C.AV_CODEC_ID_QDM2
	// AVCodecIdCook wraps AV_CODEC_ID_COOK.
	AVCodecIdCook AVCodecID = C.AV_CODEC_ID_COOK
	// AVCodecIdTruespeech wraps AV_CODEC_ID_TRUESPEECH.
	AVCodecIdTruespeech AVCodecID = C.AV_CODEC_ID_TRUESPEECH
	// AVCodecIdTta wraps AV_CODEC_ID_TTA.
	AVCodecIdTta AVCodecID = C.AV_CODEC_ID_TTA
	// AVCodecIdSmackaudio wraps AV_CODEC_ID_SMACKAUDIO.
	AVCodecIdSmackaudio AVCodecID = C.AV_CODEC_ID_SMACKAUDIO
	// AVCodecIdQcelp wraps AV_CODEC_ID_QCELP.
	AVCodecIdQcelp AVCodecID = C.AV_CODEC_ID_QCELP
	// AVCodecIdWavpack wraps AV_CODEC_ID_WAVPACK.
	AVCodecIdWavpack AVCodecID = C.AV_CODEC_ID_WAVPACK
	// AVCodecIdDsicinaudio wraps AV_CODEC_ID_DSICINAUDIO.
	AVCodecIdDsicinaudio AVCodecID = C.AV_CODEC_ID_DSICINAUDIO
	// AVCodecIdImc wraps AV_CODEC_ID_IMC.
	AVCodecIdImc AVCodecID = C.AV_CODEC_ID_IMC
	// AVCodecIdMusepack7 wraps AV_CODEC_ID_MUSEPACK7.
	AVCodecIdMusepack7 AVCodecID = C.AV_CODEC_ID_MUSEPACK7
	// AVCodecIdMlp wraps AV_CODEC_ID_MLP.
	AVCodecIdMlp AVCodecID = C.AV_CODEC_ID_MLP
	// AVCodecIdGsmMs wraps AV_CODEC_ID_GSM_MS.
	AVCodecIdGsmMs AVCodecID = C.AV_CODEC_ID_GSM_MS
	// AVCodecIdAtrac3 wraps AV_CODEC_ID_ATRAC3.
	AVCodecIdAtrac3 AVCodecID = C.AV_CODEC_ID_ATRAC3
	// AVCodecIdApe wraps AV_CODEC_ID_APE.
	AVCodecIdApe AVCodecID = C.AV_CODEC_ID_APE
	// AVCodecIdNellymoser wraps AV_CODEC_ID_NELLYMOSER.
	AVCodecIdNellymoser AVCodecID = C.AV_CODEC_ID_NELLYMOSER
	// AVCodecIdMusepack8 wraps AV_CODEC_ID_MUSEPACK8.
	AVCodecIdMusepack8 AVCodecID = C.AV_CODEC_ID_MUSEPACK8
	// AVCodecIdSpeex wraps AV_CODEC_ID_SPEEX.
	AVCodecIdSpeex AVCodecID = C.AV_CODEC_ID_SPEEX
	// AVCodecIdWmavoice wraps AV_CODEC_ID_WMAVOICE.
	AVCodecIdWmavoice AVCodecID = C.AV_CODEC_ID_WMAVOICE
	// AVCodecIdWmapro wraps AV_CODEC_ID_WMAPRO.
	AVCodecIdWmapro AVCodecID = C.AV_CODEC_ID_WMAPRO
	// AVCodecIdWmalossless wraps AV_CODEC_ID_WMALOSSLESS.
	AVCodecIdWmalossless AVCodecID = C.AV_CODEC_ID_WMALOSSLESS
	// AVCodecIdAtrac3P wraps AV_CODEC_ID_ATRAC3P.
	AVCodecIdAtrac3P AVCodecID = C.AV_CODEC_ID_ATRAC3P
	// AVCodecIdEac3 wraps AV_CODEC_ID_EAC3.
	AVCodecIdEac3 AVCodecID = C.AV_CODEC_ID_EAC3
	// AVCodecIdSipr wraps AV_CODEC_ID_SIPR.
	AVCodecIdSipr AVCodecID = C.AV_CODEC_ID_SIPR
	// AVCodecIdMp1 wraps AV_CODEC_ID_MP1.
	AVCodecIdMp1 AVCodecID = C.AV_CODEC_ID_MP1
	// AVCodecIdTwinvq wraps AV_CODEC_ID_TWINVQ.
	AVCodecIdTwinvq AVCodecID = C.AV_CODEC_ID_TWINVQ
	// AVCodecIdTruehd wraps AV_CODEC_ID_TRUEHD.
	AVCodecIdTruehd AVCodecID = C.AV_CODEC_ID_TRUEHD
	// AVCodecIdMp4Als wraps AV_CODEC_ID_MP4ALS.
	AVCodecIdMp4Als AVCodecID = C.AV_CODEC_ID_MP4ALS
	// AVCodecIdAtrac1 wraps AV_CODEC_ID_ATRAC1.
	AVCodecIdAtrac1 AVCodecID = C.AV_CODEC_ID_ATRAC1
	// AVCodecIdBinkaudioRdft wraps AV_CODEC_ID_BINKAUDIO_RDFT.
	AVCodecIdBinkaudioRdft AVCodecID = C.AV_CODEC_ID_BINKAUDIO_RDFT
	// AVCodecIdBinkaudioDct wraps AV_CODEC_ID_BINKAUDIO_DCT.
	AVCodecIdBinkaudioDct AVCodecID = C.AV_CODEC_ID_BINKAUDIO_DCT
	// AVCodecIdAacLatm wraps AV_CODEC_ID_AAC_LATM.
	AVCodecIdAacLatm AVCodecID = C.AV_CODEC_ID_AAC_LATM
	// AVCodecIdQdmc wraps AV_CODEC_ID_QDMC.
	AVCodecIdQdmc AVCodecID = C.AV_CODEC_ID_QDMC
	// AVCodecIdCelt wraps AV_CODEC_ID_CELT.
	AVCodecIdCelt AVCodecID = C.AV_CODEC_ID_CELT
	// AVCodecIdG7231 wraps AV_CODEC_ID_G723_1.
	AVCodecIdG7231 AVCodecID = C.AV_CODEC_ID_G723_1
	// AVCodecIdG729 wraps AV_CODEC_ID_G729.
	AVCodecIdG729 AVCodecID = C.AV_CODEC_ID_G729
	// AVCodecId8SvxExp wraps AV_CODEC_ID_8SVX_EXP.
	AVCodecId8SvxExp AVCodecID = C.AV_CODEC_ID_8SVX_EXP
	// AVCodecId8SvxFib wraps AV_CODEC_ID_8SVX_FIB.
	AVCodecId8SvxFib AVCodecID = C.AV_CODEC_ID_8SVX_FIB
	// AVCodecIdBmvAudio wraps AV_CODEC_ID_BMV_AUDIO.
	AVCodecIdBmvAudio AVCodecID = C.AV_CODEC_ID_BMV_AUDIO
	// AVCodecIdRalf wraps AV_CODEC_ID_RALF.
	AVCodecIdRalf AVCodecID = C.AV_CODEC_ID_RALF
	// AVCodecIdIac wraps AV_CODEC_ID_IAC.
	AVCodecIdIac AVCodecID = C.AV_CODEC_ID_IAC
	// AVCodecIdIlbc wraps AV_CODEC_ID_ILBC.
	AVCodecIdIlbc AVCodecID = C.AV_CODEC_ID_ILBC
	// AVCodecIdOpus wraps AV_CODEC_ID_OPUS.
	AVCodecIdOpus AVCodecID = C.AV_CODEC_ID_OPUS
	// AVCodecIdComfortNoise wraps AV_CODEC_ID_COMFORT_NOISE.
	AVCodecIdComfortNoise AVCodecID = C.AV_CODEC_ID_COMFORT_NOISE
	// AVCodecIdTak wraps AV_CODEC_ID_TAK.
	AVCodecIdTak AVCodecID = C.AV_CODEC_ID_TAK
	// AVCodecIdMetasound wraps AV_CODEC_ID_METASOUND.
	AVCodecIdMetasound AVCodecID = C.AV_CODEC_ID_METASOUND
	// AVCodecIdPafAudio wraps AV_CODEC_ID_PAF_AUDIO.
	AVCodecIdPafAudio AVCodecID = C.AV_CODEC_ID_PAF_AUDIO
	// AVCodecIdOn2Avc wraps AV_CODEC_ID_ON2AVC.
	AVCodecIdOn2Avc AVCodecID = C.AV_CODEC_ID_ON2AVC
	// AVCodecIdDssSp wraps AV_CODEC_ID_DSS_SP.
	AVCodecIdDssSp AVCodecID = C.AV_CODEC_ID_DSS_SP
	// AVCodecIdCodec2 wraps AV_CODEC_ID_CODEC2.
	AVCodecIdCodec2 AVCodecID = C.AV_CODEC_ID_CODEC2
	// AVCodecIdFfwavesynth wraps AV_CODEC_ID_FFWAVESYNTH.
	AVCodecIdFfwavesynth AVCodecID = C.AV_CODEC_ID_FFWAVESYNTH
	// AVCodecIdSonic wraps AV_CODEC_ID_SONIC.
	AVCodecIdSonic AVCodecID = C.AV_CODEC_ID_SONIC
	// AVCodecIdSonicLs wraps AV_CODEC_ID_SONIC_LS.
	AVCodecIdSonicLs AVCodecID = C.AV_CODEC_ID_SONIC_LS
	// AVCodecIdEvrc wraps AV_CODEC_ID_EVRC.
	AVCodecIdEvrc AVCodecID = C.AV_CODEC_ID_EVRC
	// AVCodecIdSmv wraps AV_CODEC_ID_SMV.
	AVCodecIdSmv AVCodecID = C.AV_CODEC_ID_SMV
	// AVCodecIdDsdLsbf wraps AV_CODEC_ID_DSD_LSBF.
	AVCodecIdDsdLsbf AVCodecID = C.AV_CODEC_ID_DSD_LSBF
	// AVCodecIdDsdMsbf wraps AV_CODEC_ID_DSD_MSBF.
	AVCodecIdDsdMsbf AVCodecID = C.AV_CODEC_ID_DSD_MSBF
	// AVCodecIdDsdLsbfPlanar wraps AV_CODEC_ID_DSD_LSBF_PLANAR.
	AVCodecIdDsdLsbfPlanar AVCodecID = C.AV_CODEC_ID_DSD_LSBF_PLANAR
	// AVCodecIdDsdMsbfPlanar wraps AV_CODEC_ID_DSD_MSBF_PLANAR.
	AVCodecIdDsdMsbfPlanar AVCodecID = C.AV_CODEC_ID_DSD_MSBF_PLANAR
	// AVCodecId4Gv wraps AV_CODEC_ID_4GV.
	AVCodecId4Gv AVCodecID = C.AV_CODEC_ID_4GV
	// AVCodecIdInterplayAcm wraps AV_CODEC_ID_INTERPLAY_ACM.
	AVCodecIdInterplayAcm AVCodecID = C.AV_CODEC_ID_INTERPLAY_ACM
	// AVCodecIdXma1 wraps AV_CODEC_ID_XMA1.
	AVCodecIdXma1 AVCodecID = C.AV_CODEC_ID_XMA1
	// AVCodecIdXma2 wraps AV_CODEC_ID_XMA2.
	AVCodecIdXma2 AVCodecID = C.AV_CODEC_ID_XMA2
	// AVCodecIdDst wraps AV_CODEC_ID_DST.
	AVCodecIdDst AVCodecID = C.AV_CODEC_ID_DST
	// AVCodecIdAtrac3Al wraps AV_CODEC_ID_ATRAC3AL.
	AVCodecIdAtrac3Al AVCodecID = C.AV_CODEC_ID_ATRAC3AL
	// AVCodecIdAtrac3Pal wraps AV_CODEC_ID_ATRAC3PAL.
	AVCodecIdAtrac3Pal AVCodecID = C.AV_CODEC_ID_ATRAC3PAL
	// AVCodecIdDolbyE wraps AV_CODEC_ID_DOLBY_E.
	AVCodecIdDolbyE AVCodecID = C.AV_CODEC_ID_DOLBY_E
	// AVCodecIdAptx wraps AV_CODEC_ID_APTX.
	AVCodecIdAptx AVCodecID = C.AV_CODEC_ID_APTX
	// AVCodecIdAptxHd wraps AV_CODEC_ID_APTX_HD.
	AVCodecIdAptxHd AVCodecID = C.AV_CODEC_ID_APTX_HD
	// AVCodecIdSbc wraps AV_CODEC_ID_SBC.
	AVCodecIdSbc AVCodecID = C.AV_CODEC_ID_SBC
	// AVCodecIdAtrac9 wraps AV_CODEC_ID_ATRAC9.
	AVCodecIdAtrac9 AVCodecID = C.AV_CODEC_ID_ATRAC9
	// AVCodecIdHcom wraps AV_CODEC_ID_HCOM.
	AVCodecIdHcom AVCodecID = C.AV_CODEC_ID_HCOM
	// AVCodecIdAcelpKelvin wraps AV_CODEC_ID_ACELP_KELVIN.
	AVCodecIdAcelpKelvin AVCodecID = C.AV_CODEC_ID_ACELP_KELVIN
	// AVCodecIdMpegh3DAudio wraps AV_CODEC_ID_MPEGH_3D_AUDIO.
	AVCodecIdMpegh3DAudio AVCodecID = C.AV_CODEC_ID_MPEGH_3D_AUDIO
	// AVCodecIdSiren wraps AV_CODEC_ID_SIREN.
	AVCodecIdSiren AVCodecID = C.AV_CODEC_ID_SIREN
	// AVCodecIdHca wraps AV_CODEC_ID_HCA.
	AVCodecIdHca AVCodecID = C.AV_CODEC_ID_HCA
	// AVCodecIdFastaudio wraps AV_CODEC_ID_FASTAUDIO.
	AVCodecIdFastaudio AVCodecID = C.AV_CODEC_ID_FASTAUDIO
	// AVCodecIdMsnsiren wraps AV_CODEC_ID_MSNSIREN.
	AVCodecIdMsnsiren AVCodecID = C.AV_CODEC_ID_MSNSIREN
	// AVCodecIdDfpwm wraps AV_CODEC_ID_DFPWM.
	AVCodecIdDfpwm AVCodecID = C.AV_CODEC_ID_DFPWM
	// AVCodecIdBonk wraps AV_CODEC_ID_BONK.
	AVCodecIdBonk AVCodecID = C.AV_CODEC_ID_BONK
	// AVCodecIdMisc4 wraps AV_CODEC_ID_MISC4.
	AVCodecIdMisc4 AVCodecID = C.AV_CODEC_ID_MISC4
	// AVCodecIdApac wraps AV_CODEC_ID_APAC.
	AVCodecIdApac AVCodecID = C.AV_CODEC_ID_APAC
	// AVCodecIdFtr wraps AV_CODEC_ID_FTR.
	AVCodecIdFtr AVCodecID = C.AV_CODEC_ID_FTR
	// AVCodecIdWavarc wraps AV_CODEC_ID_WAVARC.
	AVCodecIdWavarc AVCodecID = C.AV_CODEC_ID_WAVARC
	// AVCodecIdRka wraps AV_CODEC_ID_RKA.
	AVCodecIdRka AVCodecID = C.AV_CODEC_ID_RKA
	// AVCodecIdFirstSubtitle wraps AV_CODEC_ID_FIRST_SUBTITLE.
	AVCodecIdFirstSubtitle AVCodecID = C.AV_CODEC_ID_FIRST_SUBTITLE
	// AVCodecIdDvdSubtitle wraps AV_CODEC_ID_DVD_SUBTITLE.
	AVCodecIdDvdSubtitle AVCodecID = C.AV_CODEC_ID_DVD_SUBTITLE
	// AVCodecIdDvbSubtitle wraps AV_CODEC_ID_DVB_SUBTITLE.
	AVCodecIdDvbSubtitle AVCodecID = C.AV_CODEC_ID_DVB_SUBTITLE
	// AVCodecIdText wraps AV_CODEC_ID_TEXT.
	AVCodecIdText AVCodecID = C.AV_CODEC_ID_TEXT
	// AVCodecIdXsub wraps AV_CODEC_ID_XSUB.
	AVCodecIdXsub AVCodecID = C.AV_CODEC_ID_XSUB
	// AVCodecIdSsa wraps AV_CODEC_ID_SSA.
	AVCodecIdSsa AVCodecID = C.AV_CODEC_ID_SSA
	// AVCodecIdMovText wraps AV_CODEC_ID_MOV_TEXT.
	AVCodecIdMovText AVCodecID = C.AV_CODEC_ID_MOV_TEXT
	// AVCodecIdHdmvPgsSubtitle wraps AV_CODEC_ID_HDMV_PGS_SUBTITLE.
	AVCodecIdHdmvPgsSubtitle AVCodecID = C.AV_CODEC_ID_HDMV_PGS_SUBTITLE
	// AVCodecIdDvbTeletext wraps AV_CODEC_ID_DVB_TELETEXT.
	AVCodecIdDvbTeletext AVCodecID = C.AV_CODEC_ID_DVB_TELETEXT
	// AVCodecIdSrt wraps AV_CODEC_ID_SRT.
	AVCodecIdSrt AVCodecID = C.AV_CODEC_ID_SRT
	// AVCodecIdMicrodvd wraps AV_CODEC_ID_MICRODVD.
	AVCodecIdMicrodvd AVCodecID = C.AV_CODEC_ID_MICRODVD
	// AVCodecIdEia608 wraps AV_CODEC_ID_EIA_608.
	AVCodecIdEia608 AVCodecID = C.AV_CODEC_ID_EIA_608
	// AVCodecIdJacosub wraps AV_CODEC_ID_JACOSUB.
	AVCodecIdJacosub AVCodecID = C.AV_CODEC_ID_JACOSUB
	// AVCodecIdSami wraps AV_CODEC_ID_SAMI.
	AVCodecIdSami AVCodecID = C.AV_CODEC_ID_SAMI
	// AVCodecIdRealtext wraps AV_CODEC_ID_REALTEXT.
	AVCodecIdRealtext AVCodecID = C.AV_CODEC_ID_REALTEXT
	// AVCodecIdStl wraps AV_CODEC_ID_STL.
	AVCodecIdStl AVCodecID = C.AV_CODEC_ID_STL
	// AVCodecIdSubviewer1 wraps AV_CODEC_ID_SUBVIEWER1.
	AVCodecIdSubviewer1 AVCodecID = C.AV_CODEC_ID_SUBVIEWER1
	// AVCodecIdSubviewer wraps AV_CODEC_ID_SUBVIEWER.
	AVCodecIdSubviewer AVCodecID = C.AV_CODEC_ID_SUBVIEWER
	// AVCodecIdSubrip wraps AV_CODEC_ID_SUBRIP.
	AVCodecIdSubrip AVCodecID = C.AV_CODEC_ID_SUBRIP
	// AVCodecIdWebvtt wraps AV_CODEC_ID_WEBVTT.
	AVCodecIdWebvtt AVCodecID = C.AV_CODEC_ID_WEBVTT
	// AVCodecIdMpl2 wraps AV_CODEC_ID_MPL2.
	AVCodecIdMpl2 AVCodecID = C.AV_CODEC_ID_MPL2
	// AVCodecIdVplayer wraps AV_CODEC_ID_VPLAYER.
	AVCodecIdVplayer AVCodecID = C.AV_CODEC_ID_VPLAYER
	// AVCodecIdPjs wraps AV_CODEC_ID_PJS.
	AVCodecIdPjs AVCodecID = C.AV_CODEC_ID_PJS
	// AVCodecIdAss wraps AV_CODEC_ID_ASS.
	AVCodecIdAss AVCodecID = C.AV_CODEC_ID_ASS
	// AVCodecIdHdmvTextSubtitle wraps AV_CODEC_ID_HDMV_TEXT_SUBTITLE.
	AVCodecIdHdmvTextSubtitle AVCodecID = C.AV_CODEC_ID_HDMV_TEXT_SUBTITLE
	// AVCodecIdTtml wraps AV_CODEC_ID_TTML.
	AVCodecIdTtml AVCodecID = C.AV_CODEC_ID_TTML
	// AVCodecIdAribCaption wraps AV_CODEC_ID_ARIB_CAPTION.
	AVCodecIdAribCaption AVCodecID = C.AV_CODEC_ID_ARIB_CAPTION
	// AVCodecIdFirstUnknown wraps AV_CODEC_ID_FIRST_UNKNOWN.
	AVCodecIdFirstUnknown AVCodecID = C.AV_CODEC_ID_FIRST_UNKNOWN
	// AVCodecIdTtf wraps AV_CODEC_ID_TTF.
	AVCodecIdTtf AVCodecID = C.AV_CODEC_ID_TTF
	// AVCodecIdScte35 wraps AV_CODEC_ID_SCTE_35.
	AVCodecIdScte35 AVCodecID = C.AV_CODEC_ID_SCTE_35
	// AVCodecIdEpg wraps AV_CODEC_ID_EPG.
	AVCodecIdEpg AVCodecID = C.AV_CODEC_ID_EPG
	// AVCodecIdBintext wraps AV_CODEC_ID_BINTEXT.
	AVCodecIdBintext AVCodecID = C.AV_CODEC_ID_BINTEXT
	// AVCodecIdXbin wraps AV_CODEC_ID_XBIN.
	AVCodecIdXbin AVCodecID = C.AV_CODEC_ID_XBIN
	// AVCodecIdIdf wraps AV_CODEC_ID_IDF.
	AVCodecIdIdf AVCodecID = C.AV_CODEC_ID_IDF
	// AVCodecIdOtf wraps AV_CODEC_ID_OTF.
	AVCodecIdOtf AVCodecID = C.AV_CODEC_ID_OTF
	// AVCodecIdSmpteKlv wraps AV_CODEC_ID_SMPTE_KLV.
	AVCodecIdSmpteKlv AVCodecID = C.AV_CODEC_ID_SMPTE_KLV
	// AVCodecIdDvdNav wraps AV_CODEC_ID_DVD_NAV.
	AVCodecIdDvdNav AVCodecID = C.AV_CODEC_ID_DVD_NAV
	// AVCodecIdTimedId3 wraps AV_CODEC_ID_TIMED_ID3.
	AVCodecIdTimedId3 AVCodecID = C.AV_CODEC_ID_TIMED_ID3
	// AVCodecIdBinData wraps AV_CODEC_ID_BIN_DATA.
	AVCodecIdBinData AVCodecID = C.AV_CODEC_ID_BIN_DATA
	// AVCodecIdProbe wraps AV_CODEC_ID_PROBE.
	AVCodecIdProbe AVCodecID = C.AV_CODEC_ID_PROBE
	// AVCodecIdMpeg2Ts wraps AV_CODEC_ID_MPEG2TS.
	AVCodecIdMpeg2Ts AVCodecID = C.AV_CODEC_ID_MPEG2TS
	// AVCodecIdMpeg4Systems wraps AV_CODEC_ID_MPEG4SYSTEMS.
	AVCodecIdMpeg4Systems AVCodecID = C.AV_CODEC_ID_MPEG4SYSTEMS
	// AVCodecIdFfmetadata wraps AV_CODEC_ID_FFMETADATA.
	AVCodecIdFfmetadata AVCodecID = C.AV_CODEC_ID_FFMETADATA
	// AVCodecIdWrappedAvframe wraps AV_CODEC_ID_WRAPPED_AVFRAME.
	AVCodecIdWrappedAvframe AVCodecID = C.AV_CODEC_ID_WRAPPED_AVFRAME
	// AVCodecIdVnull wraps AV_CODEC_ID_VNULL.
	AVCodecIdVnull AVCodecID = C.AV_CODEC_ID_VNULL
	// AVCodecIdAnull wraps AV_CODEC_ID_ANULL.
	AVCodecIdAnull AVCodecID = C.AV_CODEC_ID_ANULL
)

// --- Enum AVFieldOrder ---

// AVFieldOrder wraps AVFieldOrder.
type AVFieldOrder C.enum_AVFieldOrder

const (
	// AVFieldUnknown wraps AV_FIELD_UNKNOWN.
	AVFieldUnknown AVFieldOrder = C.AV_FIELD_UNKNOWN
	// AVFieldProgressive wraps AV_FIELD_PROGRESSIVE.
	AVFieldProgressive AVFieldOrder = C.AV_FIELD_PROGRESSIVE
	// AVFieldTt wraps AV_FIELD_TT.
	AVFieldTt AVFieldOrder = C.AV_FIELD_TT
	// AVFieldBb wraps AV_FIELD_BB.
	AVFieldBb AVFieldOrder = C.AV_FIELD_BB
	// AVFieldTb wraps AV_FIELD_TB.
	AVFieldTb AVFieldOrder = C.AV_FIELD_TB
	// AVFieldBt wraps AV_FIELD_BT.
	AVFieldBt AVFieldOrder = C.AV_FIELD_BT
)

// --- Enum AVDiscard ---

// AVDiscard wraps AVDiscard.
type AVDiscard C.enum_AVDiscard

const (
	// AVDiscardNone wraps AVDISCARD_NONE.
	AVDiscardNone AVDiscard = C.AVDISCARD_NONE
	// AVDiscardDefault wraps AVDISCARD_DEFAULT.
	AVDiscardDefault AVDiscard = C.AVDISCARD_DEFAULT
	// AVDiscardNonref wraps AVDISCARD_NONREF.
	AVDiscardNonref AVDiscard = C.AVDISCARD_NONREF
	// AVDiscardBidir wraps AVDISCARD_BIDIR.
	AVDiscardBidir AVDiscard = C.AVDISCARD_BIDIR
	// AVDiscardNonintra wraps AVDISCARD_NONINTRA.
	AVDiscardNonintra AVDiscard = C.AVDISCARD_NONINTRA
	// AVDiscardNonkey wraps AVDISCARD_NONKEY.
	AVDiscardNonkey AVDiscard = C.AVDISCARD_NONKEY
	// AVDiscardAll wraps AVDISCARD_ALL.
	AVDiscardAll AVDiscard = C.AVDISCARD_ALL
)

// --- Enum AVAudioServiceType ---

// AVAudioServiceType wraps AVAudioServiceType.
type AVAudioServiceType C.enum_AVAudioServiceType

const (
	// AVAudioServiceTypeMain wraps AV_AUDIO_SERVICE_TYPE_MAIN.
	AVAudioServiceTypeMain AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_MAIN
	// AVAudioServiceTypeEffects wraps AV_AUDIO_SERVICE_TYPE_EFFECTS.
	AVAudioServiceTypeEffects AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EFFECTS
	// AVAudioServiceTypeVisuallyImpaired wraps AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED.
	AVAudioServiceTypeVisuallyImpaired AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED
	// AVAudioServiceTypeHearingImpaired wraps AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED.
	AVAudioServiceTypeHearingImpaired AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED
	// AVAudioServiceTypeDialogue wraps AV_AUDIO_SERVICE_TYPE_DIALOGUE.
	AVAudioServiceTypeDialogue AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_DIALOGUE
	// AVAudioServiceTypeCommentary wraps AV_AUDIO_SERVICE_TYPE_COMMENTARY.
	AVAudioServiceTypeCommentary AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_COMMENTARY
	// AVAudioServiceTypeEmergency wraps AV_AUDIO_SERVICE_TYPE_EMERGENCY.
	AVAudioServiceTypeEmergency AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EMERGENCY
	// AVAudioServiceTypeVoiceOver wraps AV_AUDIO_SERVICE_TYPE_VOICE_OVER.
	AVAudioServiceTypeVoiceOver AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VOICE_OVER
	// AVAudioServiceTypeKaraoke wraps AV_AUDIO_SERVICE_TYPE_KARAOKE.
	AVAudioServiceTypeKaraoke AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_KARAOKE
	// AVAudioServiceTypeNb wraps AV_AUDIO_SERVICE_TYPE_NB.
	AVAudioServiceTypeNb AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_NB
)

// --- Enum AVPacketSideDataType ---

// AVPacketSideDataType wraps AVPacketSideDataType.
type AVPacketSideDataType C.enum_AVPacketSideDataType

const (
	// AVPktDataPalette wraps AV_PKT_DATA_PALETTE.
	AVPktDataPalette AVPacketSideDataType = C.AV_PKT_DATA_PALETTE
	// AVPktDataNewExtradata wraps AV_PKT_DATA_NEW_EXTRADATA.
	AVPktDataNewExtradata AVPacketSideDataType = C.AV_PKT_DATA_NEW_EXTRADATA
	// AVPktDataParamChange wraps AV_PKT_DATA_PARAM_CHANGE.
	AVPktDataParamChange AVPacketSideDataType = C.AV_PKT_DATA_PARAM_CHANGE
	// AVPktDataH263MbInfo wraps AV_PKT_DATA_H263_MB_INFO.
	AVPktDataH263MbInfo AVPacketSideDataType = C.AV_PKT_DATA_H263_MB_INFO
	// AVPktDataReplaygain wraps AV_PKT_DATA_REPLAYGAIN.
	AVPktDataReplaygain AVPacketSideDataType = C.AV_PKT_DATA_REPLAYGAIN
	// AVPktDataDisplaymatrix wraps AV_PKT_DATA_DISPLAYMATRIX.
	AVPktDataDisplaymatrix AVPacketSideDataType = C.AV_PKT_DATA_DISPLAYMATRIX
	// AVPktDataStereo3D wraps AV_PKT_DATA_STEREO3D.
	AVPktDataStereo3D AVPacketSideDataType = C.AV_PKT_DATA_STEREO3D
	// AVPktDataAudioServiceType wraps AV_PKT_DATA_AUDIO_SERVICE_TYPE.
	AVPktDataAudioServiceType AVPacketSideDataType = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE
	// AVPktDataQualityStats wraps AV_PKT_DATA_QUALITY_STATS.
	AVPktDataQualityStats AVPacketSideDataType = C.AV_PKT_DATA_QUALITY_STATS
	// AVPktDataFallbackTrack wraps AV_PKT_DATA_FALLBACK_TRACK.
	AVPktDataFallbackTrack AVPacketSideDataType = C.AV_PKT_DATA_FALLBACK_TRACK
	// AVPktDataCpbProperties wraps AV_PKT_DATA_CPB_PROPERTIES.
	AVPktDataCpbProperties AVPacketSideDataType = C.AV_PKT_DATA_CPB_PROPERTIES
	// AVPktDataSkipSamples wraps AV_PKT_DATA_SKIP_SAMPLES.
	AVPktDataSkipSamples AVPacketSideDataType = C.AV_PKT_DATA_SKIP_SAMPLES
	// AVPktDataJpDualmono wraps AV_PKT_DATA_JP_DUALMONO.
	AVPktDataJpDualmono AVPacketSideDataType = C.AV_PKT_DATA_JP_DUALMONO
	// AVPktDataStringsMetadata wraps AV_PKT_DATA_STRINGS_METADATA.
	AVPktDataStringsMetadata AVPacketSideDataType = C.AV_PKT_DATA_STRINGS_METADATA
	// AVPktDataSubtitlePosition wraps AV_PKT_DATA_SUBTITLE_POSITION.
	AVPktDataSubtitlePosition AVPacketSideDataType = C.AV_PKT_DATA_SUBTITLE_POSITION
	// AVPktDataMatroskaBlockadditional wraps AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL.
	AVPktDataMatroskaBlockadditional AVPacketSideDataType = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
	// AVPktDataWebvttIdentifier wraps AV_PKT_DATA_WEBVTT_IDENTIFIER.
	AVPktDataWebvttIdentifier AVPacketSideDataType = C.AV_PKT_DATA_WEBVTT_IDENTIFIER
	// AVPktDataWebvttSettings wraps AV_PKT_DATA_WEBVTT_SETTINGS.
	AVPktDataWebvttSettings AVPacketSideDataType = C.AV_PKT_DATA_WEBVTT_SETTINGS
	// AVPktDataMetadataUpdate wraps AV_PKT_DATA_METADATA_UPDATE.
	AVPktDataMetadataUpdate AVPacketSideDataType = C.AV_PKT_DATA_METADATA_UPDATE
	// AVPktDataMpegtsStreamId wraps AV_PKT_DATA_MPEGTS_STREAM_ID.
	AVPktDataMpegtsStreamId AVPacketSideDataType = C.AV_PKT_DATA_MPEGTS_STREAM_ID
	// AVPktDataMasteringDisplayMetadata wraps AV_PKT_DATA_MASTERING_DISPLAY_METADATA.
	AVPktDataMasteringDisplayMetadata AVPacketSideDataType = C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA
	// AVPktDataSpherical wraps AV_PKT_DATA_SPHERICAL.
	AVPktDataSpherical AVPacketSideDataType = C.AV_PKT_DATA_SPHERICAL
	// AVPktDataContentLightLevel wraps AV_PKT_DATA_CONTENT_LIGHT_LEVEL.
	AVPktDataContentLightLevel AVPacketSideDataType = C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL
	// AVPktDataA53Cc wraps AV_PKT_DATA_A53_CC.
	AVPktDataA53Cc AVPacketSideDataType = C.AV_PKT_DATA_A53_CC
	// AVPktDataEncryptionInitInfo wraps AV_PKT_DATA_ENCRYPTION_INIT_INFO.
	AVPktDataEncryptionInitInfo AVPacketSideDataType = C.AV_PKT_DATA_ENCRYPTION_INIT_INFO
	// AVPktDataEncryptionInfo wraps AV_PKT_DATA_ENCRYPTION_INFO.
	AVPktDataEncryptionInfo AVPacketSideDataType = C.AV_PKT_DATA_ENCRYPTION_INFO
	// AVPktDataAfd wraps AV_PKT_DATA_AFD.
	AVPktDataAfd AVPacketSideDataType = C.AV_PKT_DATA_AFD
	// AVPktDataPrft wraps AV_PKT_DATA_PRFT.
	AVPktDataPrft AVPacketSideDataType = C.AV_PKT_DATA_PRFT
	// AVPktDataIccProfile wraps AV_PKT_DATA_ICC_PROFILE.
	AVPktDataIccProfile AVPacketSideDataType = C.AV_PKT_DATA_ICC_PROFILE
	// AVPktDataDoviConf wraps AV_PKT_DATA_DOVI_CONF.
	AVPktDataDoviConf AVPacketSideDataType = C.AV_PKT_DATA_DOVI_CONF
	// AVPktDataS12MTimecode wraps AV_PKT_DATA_S12M_TIMECODE.
	AVPktDataS12MTimecode AVPacketSideDataType = C.AV_PKT_DATA_S12M_TIMECODE
	// AVPktDataDynamicHdr10Plus wraps AV_PKT_DATA_DYNAMIC_HDR10_PLUS.
	AVPktDataDynamicHdr10Plus AVPacketSideDataType = C.AV_PKT_DATA_DYNAMIC_HDR10_PLUS
	// AVPktDataNb wraps AV_PKT_DATA_NB.
	AVPktDataNb AVPacketSideDataType = C.AV_PKT_DATA_NB
)

// --- Enum AVSideDataParamChangeFlags ---

// AVSideDataParamChangeFlags wraps AVSideDataParamChangeFlags.
type AVSideDataParamChangeFlags C.enum_AVSideDataParamChangeFlags

const (
	// AVSideDataParamChangeChannelCount wraps AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT.
	AVSideDataParamChangeChannelCount AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT
	// AVSideDataParamChangeChannelLayout wraps AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT.
	AVSideDataParamChangeChannelLayout AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT
	// AVSideDataParamChangeSampleRate wraps AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE.
	AVSideDataParamChangeSampleRate AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE
	// AVSideDataParamChangeDimensions wraps AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS.
	AVSideDataParamChangeDimensions AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS
)

// --- Enum AVStreamParseType ---

// AVStreamParseType wraps AVStreamParseType.
type AVStreamParseType C.enum_AVStreamParseType

const (
	// AVStreamParseNone wraps AVSTREAM_PARSE_NONE.
	AVStreamParseNone AVStreamParseType = C.AVSTREAM_PARSE_NONE
	// AVStreamParseFull wraps AVSTREAM_PARSE_FULL.
	AVStreamParseFull AVStreamParseType = C.AVSTREAM_PARSE_FULL
	// AVStreamParseHeaders wraps AVSTREAM_PARSE_HEADERS.
	AVStreamParseHeaders AVStreamParseType = C.AVSTREAM_PARSE_HEADERS
	// AVStreamParseTimestamps wraps AVSTREAM_PARSE_TIMESTAMPS.
	AVStreamParseTimestamps AVStreamParseType = C.AVSTREAM_PARSE_TIMESTAMPS
	// AVStreamParseFullOnce wraps AVSTREAM_PARSE_FULL_ONCE.
	AVStreamParseFullOnce AVStreamParseType = C.AVSTREAM_PARSE_FULL_ONCE
	// AVStreamParseFullRaw wraps AVSTREAM_PARSE_FULL_RAW.
	AVStreamParseFullRaw AVStreamParseType = C.AVSTREAM_PARSE_FULL_RAW
)

// --- Enum AVDurationEstimationMethod ---

// AVDurationEstimationMethod wraps AVDurationEstimationMethod.
type AVDurationEstimationMethod C.enum_AVDurationEstimationMethod

const (
	// AVFmtDurationFromPts wraps AVFMT_DURATION_FROM_PTS.
	AVFmtDurationFromPts AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_PTS
	// AVFmtDurationFromStream wraps AVFMT_DURATION_FROM_STREAM.
	AVFmtDurationFromStream AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_STREAM
	// AVFmtDurationFromBitrate wraps AVFMT_DURATION_FROM_BITRATE.
	AVFmtDurationFromBitrate AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_BITRATE
)

// --- Enum AVTimebaseSource ---

// AVTimebaseSource wraps AVTimebaseSource.
type AVTimebaseSource C.enum_AVTimebaseSource

const (
	// AVFmtTbcfAuto wraps AVFMT_TBCF_AUTO.
	AVFmtTbcfAuto AVTimebaseSource = C.AVFMT_TBCF_AUTO
	// AVFmtTbcfDecoder wraps AVFMT_TBCF_DECODER.
	AVFmtTbcfDecoder AVTimebaseSource = C.AVFMT_TBCF_DECODER
	// AVFmtTbcfDemuxer wraps AVFMT_TBCF_DEMUXER.
	AVFmtTbcfDemuxer AVTimebaseSource = C.AVFMT_TBCF_DEMUXER
	// AVFmtTbcfRFramerate wraps AVFMT_TBCF_R_FRAMERATE.
	AVFmtTbcfRFramerate AVTimebaseSource = C.AVFMT_TBCF_R_FRAMERATE
)

// --- Enum AVIODirEntryType ---

// AVIODirEntryType wraps AVIODirEntryType.
type AVIODirEntryType C.enum_AVIODirEntryType

const (
	// AVIOEntryUnknown wraps AVIO_ENTRY_UNKNOWN.
	AVIOEntryUnknown AVIODirEntryType = C.AVIO_ENTRY_UNKNOWN
	// AVIOEntryBlockDevice wraps AVIO_ENTRY_BLOCK_DEVICE.
	AVIOEntryBlockDevice AVIODirEntryType = C.AVIO_ENTRY_BLOCK_DEVICE
	// AVIOEntryCharacterDevice wraps AVIO_ENTRY_CHARACTER_DEVICE.
	AVIOEntryCharacterDevice AVIODirEntryType = C.AVIO_ENTRY_CHARACTER_DEVICE
	// AVIOEntryDirectory wraps AVIO_ENTRY_DIRECTORY.
	AVIOEntryDirectory AVIODirEntryType = C.AVIO_ENTRY_DIRECTORY
	// AVIOEntryNamedPipe wraps AVIO_ENTRY_NAMED_PIPE.
	AVIOEntryNamedPipe AVIODirEntryType = C.AVIO_ENTRY_NAMED_PIPE
	// AVIOEntrySymbolicLink wraps AVIO_ENTRY_SYMBOLIC_LINK.
	AVIOEntrySymbolicLink AVIODirEntryType = C.AVIO_ENTRY_SYMBOLIC_LINK
	// AVIOEntrySocket wraps AVIO_ENTRY_SOCKET.
	AVIOEntrySocket AVIODirEntryType = C.AVIO_ENTRY_SOCKET
	// AVIOEntryFile wraps AVIO_ENTRY_FILE.
	AVIOEntryFile AVIODirEntryType = C.AVIO_ENTRY_FILE
	// AVIOEntryServer wraps AVIO_ENTRY_SERVER.
	AVIOEntryServer AVIODirEntryType = C.AVIO_ENTRY_SERVER
	// AVIOEntryShare wraps AVIO_ENTRY_SHARE.
	AVIOEntryShare AVIODirEntryType = C.AVIO_ENTRY_SHARE
	// AVIOEntryWorkgroup wraps AVIO_ENTRY_WORKGROUP.
	AVIOEntryWorkgroup AVIODirEntryType = C.AVIO_ENTRY_WORKGROUP
)

// --- Enum AVIODataMarkerType ---

// AVIODataMarkerType wraps AVIODataMarkerType.
type AVIODataMarkerType C.enum_AVIODataMarkerType

const (
	// AVIODataMarkerHeader wraps AVIO_DATA_MARKER_HEADER.
	AVIODataMarkerHeader AVIODataMarkerType = C.AVIO_DATA_MARKER_HEADER
	// AVIODataMarkerSyncPoint wraps AVIO_DATA_MARKER_SYNC_POINT.
	AVIODataMarkerSyncPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_SYNC_POINT
	// AVIODataMarkerBoundaryPoint wraps AVIO_DATA_MARKER_BOUNDARY_POINT.
	AVIODataMarkerBoundaryPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_BOUNDARY_POINT
	// AVIODataMarkerUnknown wraps AVIO_DATA_MARKER_UNKNOWN.
	AVIODataMarkerUnknown AVIODataMarkerType = C.AVIO_DATA_MARKER_UNKNOWN
	// AVIODataMarkerTrailer wraps AVIO_DATA_MARKER_TRAILER.
	AVIODataMarkerTrailer AVIODataMarkerType = C.AVIO_DATA_MARKER_TRAILER
	// AVIODataMarkerFlushPoint wraps AVIO_DATA_MARKER_FLUSH_POINT.
	AVIODataMarkerFlushPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_FLUSH_POINT
)

// --- Enum AVMediaType ---

// AVMediaType wraps AVMediaType.
type AVMediaType C.enum_AVMediaType

const (
	// AVMediaTypeUnknown wraps AVMEDIA_TYPE_UNKNOWN.
	AVMediaTypeUnknown AVMediaType = C.AVMEDIA_TYPE_UNKNOWN
	// AVMediaTypeVideo wraps AVMEDIA_TYPE_VIDEO.
	AVMediaTypeVideo AVMediaType = C.AVMEDIA_TYPE_VIDEO
	// AVMediaTypeAudio wraps AVMEDIA_TYPE_AUDIO.
	AVMediaTypeAudio AVMediaType = C.AVMEDIA_TYPE_AUDIO
	// AVMediaTypeData wraps AVMEDIA_TYPE_DATA.
	AVMediaTypeData AVMediaType = C.AVMEDIA_TYPE_DATA
	// AVMediaTypeSubtitle wraps AVMEDIA_TYPE_SUBTITLE.
	AVMediaTypeSubtitle AVMediaType = C.AVMEDIA_TYPE_SUBTITLE
	// AVMediaTypeAttachment wraps AVMEDIA_TYPE_ATTACHMENT.
	AVMediaTypeAttachment AVMediaType = C.AVMEDIA_TYPE_ATTACHMENT
	// AVMediaTypeNb wraps AVMEDIA_TYPE_NB.
	AVMediaTypeNb AVMediaType = C.AVMEDIA_TYPE_NB
)

// --- Enum AVPictureType ---

// AVPictureType wraps AVPictureType.
type AVPictureType C.enum_AVPictureType

const (
	// AVPictureTypeNone wraps AV_PICTURE_TYPE_NONE.
	AVPictureTypeNone AVPictureType = C.AV_PICTURE_TYPE_NONE
	// AVPictureTypeI wraps AV_PICTURE_TYPE_I.
	AVPictureTypeI AVPictureType = C.AV_PICTURE_TYPE_I
	// AVPictureTypeP wraps AV_PICTURE_TYPE_P.
	AVPictureTypeP AVPictureType = C.AV_PICTURE_TYPE_P
	// AVPictureTypeB wraps AV_PICTURE_TYPE_B.
	AVPictureTypeB AVPictureType = C.AV_PICTURE_TYPE_B
	// AVPictureTypeS wraps AV_PICTURE_TYPE_S.
	AVPictureTypeS AVPictureType = C.AV_PICTURE_TYPE_S
	// AVPictureTypeSi wraps AV_PICTURE_TYPE_SI.
	AVPictureTypeSi AVPictureType = C.AV_PICTURE_TYPE_SI
	// AVPictureTypeSp wraps AV_PICTURE_TYPE_SP.
	AVPictureTypeSp AVPictureType = C.AV_PICTURE_TYPE_SP
	// AVPictureTypeBi wraps AV_PICTURE_TYPE_BI.
	AVPictureTypeBi AVPictureType = C.AV_PICTURE_TYPE_BI
)

// --- Enum AVChannel ---

// AVChannel wraps AVChannel.
type AVChannel C.enum_AVChannel

const (
	// AVChanNone wraps AV_CHAN_NONE.
	AVChanNone AVChannel = C.AV_CHAN_NONE
	// AVChanFrontLeft wraps AV_CHAN_FRONT_LEFT.
	AVChanFrontLeft AVChannel = C.AV_CHAN_FRONT_LEFT
	// AVChanFrontRight wraps AV_CHAN_FRONT_RIGHT.
	AVChanFrontRight AVChannel = C.AV_CHAN_FRONT_RIGHT
	// AVChanFrontCenter wraps AV_CHAN_FRONT_CENTER.
	AVChanFrontCenter AVChannel = C.AV_CHAN_FRONT_CENTER
	// AVChanLowFrequency wraps AV_CHAN_LOW_FREQUENCY.
	AVChanLowFrequency AVChannel = C.AV_CHAN_LOW_FREQUENCY
	// AVChanBackLeft wraps AV_CHAN_BACK_LEFT.
	AVChanBackLeft AVChannel = C.AV_CHAN_BACK_LEFT
	// AVChanBackRight wraps AV_CHAN_BACK_RIGHT.
	AVChanBackRight AVChannel = C.AV_CHAN_BACK_RIGHT
	// AVChanFrontLeftOfCenter wraps AV_CHAN_FRONT_LEFT_OF_CENTER.
	AVChanFrontLeftOfCenter AVChannel = C.AV_CHAN_FRONT_LEFT_OF_CENTER
	// AVChanFrontRightOfCenter wraps AV_CHAN_FRONT_RIGHT_OF_CENTER.
	AVChanFrontRightOfCenter AVChannel = C.AV_CHAN_FRONT_RIGHT_OF_CENTER
	// AVChanBackCenter wraps AV_CHAN_BACK_CENTER.
	AVChanBackCenter AVChannel = C.AV_CHAN_BACK_CENTER
	// AVChanSideLeft wraps AV_CHAN_SIDE_LEFT.
	AVChanSideLeft AVChannel = C.AV_CHAN_SIDE_LEFT
	// AVChanSideRight wraps AV_CHAN_SIDE_RIGHT.
	AVChanSideRight AVChannel = C.AV_CHAN_SIDE_RIGHT
	// AVChanTopCenter wraps AV_CHAN_TOP_CENTER.
	AVChanTopCenter AVChannel = C.AV_CHAN_TOP_CENTER
	// AVChanTopFrontLeft wraps AV_CHAN_TOP_FRONT_LEFT.
	AVChanTopFrontLeft AVChannel = C.AV_CHAN_TOP_FRONT_LEFT
	// AVChanTopFrontCenter wraps AV_CHAN_TOP_FRONT_CENTER.
	AVChanTopFrontCenter AVChannel = C.AV_CHAN_TOP_FRONT_CENTER
	// AVChanTopFrontRight wraps AV_CHAN_TOP_FRONT_RIGHT.
	AVChanTopFrontRight AVChannel = C.AV_CHAN_TOP_FRONT_RIGHT
	// AVChanTopBackLeft wraps AV_CHAN_TOP_BACK_LEFT.
	AVChanTopBackLeft AVChannel = C.AV_CHAN_TOP_BACK_LEFT
	// AVChanTopBackCenter wraps AV_CHAN_TOP_BACK_CENTER.
	AVChanTopBackCenter AVChannel = C.AV_CHAN_TOP_BACK_CENTER
	// AVChanTopBackRight wraps AV_CHAN_TOP_BACK_RIGHT.
	AVChanTopBackRight AVChannel = C.AV_CHAN_TOP_BACK_RIGHT
	// AVChanStereoLeft wraps AV_CHAN_STEREO_LEFT.
	AVChanStereoLeft AVChannel = C.AV_CHAN_STEREO_LEFT
	// AVChanStereoRight wraps AV_CHAN_STEREO_RIGHT.
	AVChanStereoRight AVChannel = C.AV_CHAN_STEREO_RIGHT
	// AVChanWideLeft wraps AV_CHAN_WIDE_LEFT.
	AVChanWideLeft AVChannel = C.AV_CHAN_WIDE_LEFT
	// AVChanWideRight wraps AV_CHAN_WIDE_RIGHT.
	AVChanWideRight AVChannel = C.AV_CHAN_WIDE_RIGHT
	// AVChanSurroundDirectLeft wraps AV_CHAN_SURROUND_DIRECT_LEFT.
	AVChanSurroundDirectLeft AVChannel = C.AV_CHAN_SURROUND_DIRECT_LEFT
	// AVChanSurroundDirectRight wraps AV_CHAN_SURROUND_DIRECT_RIGHT.
	AVChanSurroundDirectRight AVChannel = C.AV_CHAN_SURROUND_DIRECT_RIGHT
	// AVChanLowFrequency2 wraps AV_CHAN_LOW_FREQUENCY_2.
	AVChanLowFrequency2 AVChannel = C.AV_CHAN_LOW_FREQUENCY_2
	// AVChanTopSideLeft wraps AV_CHAN_TOP_SIDE_LEFT.
	AVChanTopSideLeft AVChannel = C.AV_CHAN_TOP_SIDE_LEFT
	// AVChanTopSideRight wraps AV_CHAN_TOP_SIDE_RIGHT.
	AVChanTopSideRight AVChannel = C.AV_CHAN_TOP_SIDE_RIGHT
	// AVChanBottomFrontCenter wraps AV_CHAN_BOTTOM_FRONT_CENTER.
	AVChanBottomFrontCenter AVChannel = C.AV_CHAN_BOTTOM_FRONT_CENTER
	// AVChanBottomFrontLeft wraps AV_CHAN_BOTTOM_FRONT_LEFT.
	AVChanBottomFrontLeft AVChannel = C.AV_CHAN_BOTTOM_FRONT_LEFT
	// AVChanBottomFrontRight wraps AV_CHAN_BOTTOM_FRONT_RIGHT.
	AVChanBottomFrontRight AVChannel = C.AV_CHAN_BOTTOM_FRONT_RIGHT
	// AVChanUnused wraps AV_CHAN_UNUSED.
	AVChanUnused AVChannel = C.AV_CHAN_UNUSED
	// AVChanUnknown wraps AV_CHAN_UNKNOWN.
	AVChanUnknown AVChannel = C.AV_CHAN_UNKNOWN
	// AVChanAmbisonicBase wraps AV_CHAN_AMBISONIC_BASE.
	AVChanAmbisonicBase AVChannel = C.AV_CHAN_AMBISONIC_BASE
	// AVChanAmbisonicEnd wraps AV_CHAN_AMBISONIC_END.
	AVChanAmbisonicEnd AVChannel = C.AV_CHAN_AMBISONIC_END
)

// --- Enum AVChannelOrder ---

// AVChannelOrder wraps AVChannelOrder.
type AVChannelOrder C.enum_AVChannelOrder

const (
	// AVChannelOrderUnspec wraps AV_CHANNEL_ORDER_UNSPEC.
	AVChannelOrderUnspec AVChannelOrder = C.AV_CHANNEL_ORDER_UNSPEC
	// AVChannelOrderNative wraps AV_CHANNEL_ORDER_NATIVE.
	AVChannelOrderNative AVChannelOrder = C.AV_CHANNEL_ORDER_NATIVE
	// AVChannelOrderCustom wraps AV_CHANNEL_ORDER_CUSTOM.
	AVChannelOrderCustom AVChannelOrder = C.AV_CHANNEL_ORDER_CUSTOM
	// AVChannelOrderAmbisonic wraps AV_CHANNEL_ORDER_AMBISONIC.
	AVChannelOrderAmbisonic AVChannelOrder = C.AV_CHANNEL_ORDER_AMBISONIC
)

// --- Enum AVMatrixEncoding ---

// AVMatrixEncoding wraps AVMatrixEncoding.
type AVMatrixEncoding C.enum_AVMatrixEncoding

const (
	// AVMatrixEncodingNone wraps AV_MATRIX_ENCODING_NONE.
	AVMatrixEncodingNone AVMatrixEncoding = C.AV_MATRIX_ENCODING_NONE
	// AVMatrixEncodingDolby wraps AV_MATRIX_ENCODING_DOLBY.
	AVMatrixEncodingDolby AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBY
	// AVMatrixEncodingDplii wraps AV_MATRIX_ENCODING_DPLII.
	AVMatrixEncodingDplii AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLII
	// AVMatrixEncodingDpliix wraps AV_MATRIX_ENCODING_DPLIIX.
	AVMatrixEncodingDpliix AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLIIX
	// AVMatrixEncodingDpliiz wraps AV_MATRIX_ENCODING_DPLIIZ.
	AVMatrixEncodingDpliiz AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLIIZ
	// AVMatrixEncodingDolbyex wraps AV_MATRIX_ENCODING_DOLBYEX.
	AVMatrixEncodingDolbyex AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBYEX
	// AVMatrixEncodingDolbyheadphone wraps AV_MATRIX_ENCODING_DOLBYHEADPHONE.
	AVMatrixEncodingDolbyheadphone AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBYHEADPHONE
	// AVMatrixEncodingNb wraps AV_MATRIX_ENCODING_NB.
	AVMatrixEncodingNb AVMatrixEncoding = C.AV_MATRIX_ENCODING_NB
)

// --- Enum AVFrameSideDataType ---

// AVFrameSideDataType wraps AVFrameSideDataType.
type AVFrameSideDataType C.enum_AVFrameSideDataType

const (
	// AVFrameDataPanscan wraps AV_FRAME_DATA_PANSCAN.
	AVFrameDataPanscan AVFrameSideDataType = C.AV_FRAME_DATA_PANSCAN
	// AVFrameDataA53Cc wraps AV_FRAME_DATA_A53_CC.
	AVFrameDataA53Cc AVFrameSideDataType = C.AV_FRAME_DATA_A53_CC
	// AVFrameDataStereo3D wraps AV_FRAME_DATA_STEREO3D.
	AVFrameDataStereo3D AVFrameSideDataType = C.AV_FRAME_DATA_STEREO3D
	// AVFrameDataMatrixencoding wraps AV_FRAME_DATA_MATRIXENCODING.
	AVFrameDataMatrixencoding AVFrameSideDataType = C.AV_FRAME_DATA_MATRIXENCODING
	// AVFrameDataDownmixInfo wraps AV_FRAME_DATA_DOWNMIX_INFO.
	AVFrameDataDownmixInfo AVFrameSideDataType = C.AV_FRAME_DATA_DOWNMIX_INFO
	// AVFrameDataReplaygain wraps AV_FRAME_DATA_REPLAYGAIN.
	AVFrameDataReplaygain AVFrameSideDataType = C.AV_FRAME_DATA_REPLAYGAIN
	// AVFrameDataDisplaymatrix wraps AV_FRAME_DATA_DISPLAYMATRIX.
	AVFrameDataDisplaymatrix AVFrameSideDataType = C.AV_FRAME_DATA_DISPLAYMATRIX
	// AVFrameDataAfd wraps AV_FRAME_DATA_AFD.
	AVFrameDataAfd AVFrameSideDataType = C.AV_FRAME_DATA_AFD
	// AVFrameDataMotionVectors wraps AV_FRAME_DATA_MOTION_VECTORS.
	AVFrameDataMotionVectors AVFrameSideDataType = C.AV_FRAME_DATA_MOTION_VECTORS
	// AVFrameDataSkipSamples wraps AV_FRAME_DATA_SKIP_SAMPLES.
	AVFrameDataSkipSamples AVFrameSideDataType = C.AV_FRAME_DATA_SKIP_SAMPLES
	// AVFrameDataAudioServiceType wraps AV_FRAME_DATA_AUDIO_SERVICE_TYPE.
	AVFrameDataAudioServiceType AVFrameSideDataType = C.AV_FRAME_DATA_AUDIO_SERVICE_TYPE
	// AVFrameDataMasteringDisplayMetadata wraps AV_FRAME_DATA_MASTERING_DISPLAY_METADATA.
	AVFrameDataMasteringDisplayMetadata AVFrameSideDataType = C.AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
	// AVFrameDataGopTimecode wraps AV_FRAME_DATA_GOP_TIMECODE.
	AVFrameDataGopTimecode AVFrameSideDataType = C.AV_FRAME_DATA_GOP_TIMECODE
	// AVFrameDataSpherical wraps AV_FRAME_DATA_SPHERICAL.
	AVFrameDataSpherical AVFrameSideDataType = C.AV_FRAME_DATA_SPHERICAL
	// AVFrameDataContentLightLevel wraps AV_FRAME_DATA_CONTENT_LIGHT_LEVEL.
	AVFrameDataContentLightLevel AVFrameSideDataType = C.AV_FRAME_DATA_CONTENT_LIGHT_LEVEL
	// AVFrameDataIccProfile wraps AV_FRAME_DATA_ICC_PROFILE.
	AVFrameDataIccProfile AVFrameSideDataType = C.AV_FRAME_DATA_ICC_PROFILE
	// AVFrameDataS12MTimecode wraps AV_FRAME_DATA_S12M_TIMECODE.
	AVFrameDataS12MTimecode AVFrameSideDataType = C.AV_FRAME_DATA_S12M_TIMECODE
	// AVFrameDataDynamicHdrPlus wraps AV_FRAME_DATA_DYNAMIC_HDR_PLUS.
	AVFrameDataDynamicHdrPlus AVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_PLUS
	// AVFrameDataRegionsOfInterest wraps AV_FRAME_DATA_REGIONS_OF_INTEREST.
	AVFrameDataRegionsOfInterest AVFrameSideDataType = C.AV_FRAME_DATA_REGIONS_OF_INTEREST
	// AVFrameDataVideoEncParams wraps AV_FRAME_DATA_VIDEO_ENC_PARAMS.
	AVFrameDataVideoEncParams AVFrameSideDataType = C.AV_FRAME_DATA_VIDEO_ENC_PARAMS
	// AVFrameDataSeiUnregistered wraps AV_FRAME_DATA_SEI_UNREGISTERED.
	AVFrameDataSeiUnregistered AVFrameSideDataType = C.AV_FRAME_DATA_SEI_UNREGISTERED
	// AVFrameDataFilmGrainParams wraps AV_FRAME_DATA_FILM_GRAIN_PARAMS.
	AVFrameDataFilmGrainParams AVFrameSideDataType = C.AV_FRAME_DATA_FILM_GRAIN_PARAMS
	// AVFrameDataDetectionBboxes wraps AV_FRAME_DATA_DETECTION_BBOXES.
	AVFrameDataDetectionBboxes AVFrameSideDataType = C.AV_FRAME_DATA_DETECTION_BBOXES
	// AVFrameDataDoviRpuBuffer wraps AV_FRAME_DATA_DOVI_RPU_BUFFER.
	AVFrameDataDoviRpuBuffer AVFrameSideDataType = C.AV_FRAME_DATA_DOVI_RPU_BUFFER
	// AVFrameDataDoviMetadata wraps AV_FRAME_DATA_DOVI_METADATA.
	AVFrameDataDoviMetadata AVFrameSideDataType = C.AV_FRAME_DATA_DOVI_METADATA
	// AVFrameDataDynamicHdrVivid wraps AV_FRAME_DATA_DYNAMIC_HDR_VIVID.
	AVFrameDataDynamicHdrVivid AVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_VIVID
	// AVFrameDataAmbientViewingEnvironment wraps AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT.
	AVFrameDataAmbientViewingEnvironment AVFrameSideDataType = C.AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT
)

// --- Enum AVActiveFormatDescription ---

// AVActiveFormatDescription wraps AVActiveFormatDescription.
type AVActiveFormatDescription C.enum_AVActiveFormatDescription

const (
	// AVAfdSame wraps AV_AFD_SAME.
	AVAfdSame AVActiveFormatDescription = C.AV_AFD_SAME
	// AVAfd43 wraps AV_AFD_4_3.
	AVAfd43 AVActiveFormatDescription = C.AV_AFD_4_3
	// AVAfd169 wraps AV_AFD_16_9.
	AVAfd169 AVActiveFormatDescription = C.AV_AFD_16_9
	// AVAfd149 wraps AV_AFD_14_9.
	AVAfd149 AVActiveFormatDescription = C.AV_AFD_14_9
	// AVAfd43Sp149 wraps AV_AFD_4_3_SP_14_9.
	AVAfd43Sp149 AVActiveFormatDescription = C.AV_AFD_4_3_SP_14_9
	// AVAfd169Sp149 wraps AV_AFD_16_9_SP_14_9.
	AVAfd169Sp149 AVActiveFormatDescription = C.AV_AFD_16_9_SP_14_9
	// AVAfdSp43 wraps AV_AFD_SP_4_3.
	AVAfdSp43 AVActiveFormatDescription = C.AV_AFD_SP_4_3
)

// --- Enum AVHWDeviceType ---

// AVHWDeviceType wraps AVHWDeviceType.
type AVHWDeviceType C.enum_AVHWDeviceType

const (
	// AVHwdeviceTypeNone wraps AV_HWDEVICE_TYPE_NONE.
	AVHwdeviceTypeNone AVHWDeviceType = C.AV_HWDEVICE_TYPE_NONE
	// AVHwdeviceTypeVdpau wraps AV_HWDEVICE_TYPE_VDPAU.
	AVHwdeviceTypeVdpau AVHWDeviceType = C.AV_HWDEVICE_TYPE_VDPAU
	// AVHwdeviceTypeCuda wraps AV_HWDEVICE_TYPE_CUDA.
	AVHwdeviceTypeCuda AVHWDeviceType = C.AV_HWDEVICE_TYPE_CUDA
	// AVHwdeviceTypeVaapi wraps AV_HWDEVICE_TYPE_VAAPI.
	AVHwdeviceTypeVaapi AVHWDeviceType = C.AV_HWDEVICE_TYPE_VAAPI
	// AVHwdeviceTypeDxva2 wraps AV_HWDEVICE_TYPE_DXVA2.
	AVHwdeviceTypeDxva2 AVHWDeviceType = C.AV_HWDEVICE_TYPE_DXVA2
	// AVHwdeviceTypeQsv wraps AV_HWDEVICE_TYPE_QSV.
	AVHwdeviceTypeQsv AVHWDeviceType = C.AV_HWDEVICE_TYPE_QSV
	// AVHwdeviceTypeVideotoolbox wraps AV_HWDEVICE_TYPE_VIDEOTOOLBOX.
	AVHwdeviceTypeVideotoolbox AVHWDeviceType = C.AV_HWDEVICE_TYPE_VIDEOTOOLBOX
	// AVHwdeviceTypeD3D11Va wraps AV_HWDEVICE_TYPE_D3D11VA.
	AVHwdeviceTypeD3D11Va AVHWDeviceType = C.AV_HWDEVICE_TYPE_D3D11VA
	// AVHwdeviceTypeDrm wraps AV_HWDEVICE_TYPE_DRM.
	AVHwdeviceTypeDrm AVHWDeviceType = C.AV_HWDEVICE_TYPE_DRM
	// AVHwdeviceTypeOpencl wraps AV_HWDEVICE_TYPE_OPENCL.
	AVHwdeviceTypeOpencl AVHWDeviceType = C.AV_HWDEVICE_TYPE_OPENCL
	// AVHwdeviceTypeMediacodec wraps AV_HWDEVICE_TYPE_MEDIACODEC.
	AVHwdeviceTypeMediacodec AVHWDeviceType = C.AV_HWDEVICE_TYPE_MEDIACODEC
	// AVHwdeviceTypeVulkan wraps AV_HWDEVICE_TYPE_VULKAN.
	AVHwdeviceTypeVulkan AVHWDeviceType = C.AV_HWDEVICE_TYPE_VULKAN
)

// --- Enum AVHWFrameTransferDirection ---

// AVHWFrameTransferDirection wraps AVHWFrameTransferDirection.
type AVHWFrameTransferDirection C.enum_AVHWFrameTransferDirection

const (
	// AVHwframeTransferDirectionFrom wraps AV_HWFRAME_TRANSFER_DIRECTION_FROM.
	AVHwframeTransferDirectionFrom AVHWFrameTransferDirection = C.AV_HWFRAME_TRANSFER_DIRECTION_FROM
	// AVHwframeTransferDirectionTo wraps AV_HWFRAME_TRANSFER_DIRECTION_TO.
	AVHwframeTransferDirectionTo AVHWFrameTransferDirection = C.AV_HWFRAME_TRANSFER_DIRECTION_TO
)

// --- Enum AVClassCategory ---

// AVClassCategory wraps AVClassCategory.
type AVClassCategory C.AVClassCategory

const (
	// AVClassCategoryNa wraps AV_CLASS_CATEGORY_NA.
	AVClassCategoryNa AVClassCategory = C.AV_CLASS_CATEGORY_NA
	// AVClassCategoryInput wraps AV_CLASS_CATEGORY_INPUT.
	AVClassCategoryInput AVClassCategory = C.AV_CLASS_CATEGORY_INPUT
	// AVClassCategoryOutput wraps AV_CLASS_CATEGORY_OUTPUT.
	AVClassCategoryOutput AVClassCategory = C.AV_CLASS_CATEGORY_OUTPUT
	// AVClassCategoryMuxer wraps AV_CLASS_CATEGORY_MUXER.
	AVClassCategoryMuxer AVClassCategory = C.AV_CLASS_CATEGORY_MUXER
	// AVClassCategoryDemuxer wraps AV_CLASS_CATEGORY_DEMUXER.
	AVClassCategoryDemuxer AVClassCategory = C.AV_CLASS_CATEGORY_DEMUXER
	// AVClassCategoryEncoder wraps AV_CLASS_CATEGORY_ENCODER.
	AVClassCategoryEncoder AVClassCategory = C.AV_CLASS_CATEGORY_ENCODER
	// AVClassCategoryDecoder wraps AV_CLASS_CATEGORY_DECODER.
	AVClassCategoryDecoder AVClassCategory = C.AV_CLASS_CATEGORY_DECODER
	// AVClassCategoryFilter wraps AV_CLASS_CATEGORY_FILTER.
	AVClassCategoryFilter AVClassCategory = C.AV_CLASS_CATEGORY_FILTER
	// AVClassCategoryBitstreamFilter wraps AV_CLASS_CATEGORY_BITSTREAM_FILTER.
	AVClassCategoryBitstreamFilter AVClassCategory = C.AV_CLASS_CATEGORY_BITSTREAM_FILTER
	// AVClassCategorySwscaler wraps AV_CLASS_CATEGORY_SWSCALER.
	AVClassCategorySwscaler AVClassCategory = C.AV_CLASS_CATEGORY_SWSCALER
	// AVClassCategorySwresampler wraps AV_CLASS_CATEGORY_SWRESAMPLER.
	AVClassCategorySwresampler AVClassCategory = C.AV_CLASS_CATEGORY_SWRESAMPLER
	// AVClassCategoryDeviceVideoOutput wraps AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT.
	AVClassCategoryDeviceVideoOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT
	// AVClassCategoryDeviceVideoInput wraps AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT.
	AVClassCategoryDeviceVideoInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT
	// AVClassCategoryDeviceAudioOutput wraps AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT.
	AVClassCategoryDeviceAudioOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT
	// AVClassCategoryDeviceAudioInput wraps AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT.
	AVClassCategoryDeviceAudioInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT
	// AVClassCategoryDeviceOutput wraps AV_CLASS_CATEGORY_DEVICE_OUTPUT.
	AVClassCategoryDeviceOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_OUTPUT
	// AVClassCategoryDeviceInput wraps AV_CLASS_CATEGORY_DEVICE_INPUT.
	AVClassCategoryDeviceInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_INPUT
	// AVClassCategoryNb wraps AV_CLASS_CATEGORY_NB.
	AVClassCategoryNb AVClassCategory = C.AV_CLASS_CATEGORY_NB
)

// --- Enum AVOptionType ---

// AVOptionType wraps AVOptionType.
type AVOptionType C.enum_AVOptionType

const (
	// AVOptTypeFlags wraps AV_OPT_TYPE_FLAGS.
	AVOptTypeFlags AVOptionType = C.AV_OPT_TYPE_FLAGS
	// AVOptTypeInt wraps AV_OPT_TYPE_INT.
	AVOptTypeInt AVOptionType = C.AV_OPT_TYPE_INT
	// AVOptTypeInt64 wraps AV_OPT_TYPE_INT64.
	AVOptTypeInt64 AVOptionType = C.AV_OPT_TYPE_INT64
	// AVOptTypeDouble wraps AV_OPT_TYPE_DOUBLE.
	AVOptTypeDouble AVOptionType = C.AV_OPT_TYPE_DOUBLE
	// AVOptTypeFloat wraps AV_OPT_TYPE_FLOAT.
	AVOptTypeFloat AVOptionType = C.AV_OPT_TYPE_FLOAT
	// AVOptTypeString wraps AV_OPT_TYPE_STRING.
	AVOptTypeString AVOptionType = C.AV_OPT_TYPE_STRING
	// AVOptTypeRational wraps AV_OPT_TYPE_RATIONAL.
	AVOptTypeRational AVOptionType = C.AV_OPT_TYPE_RATIONAL
	// AVOptTypeBinary wraps AV_OPT_TYPE_BINARY.
	AVOptTypeBinary AVOptionType = C.AV_OPT_TYPE_BINARY
	// AVOptTypeDict wraps AV_OPT_TYPE_DICT.
	AVOptTypeDict AVOptionType = C.AV_OPT_TYPE_DICT
	// AVOptTypeUint64 wraps AV_OPT_TYPE_UINT64.
	AVOptTypeUint64 AVOptionType = C.AV_OPT_TYPE_UINT64
	// AVOptTypeConst wraps AV_OPT_TYPE_CONST.
	AVOptTypeConst AVOptionType = C.AV_OPT_TYPE_CONST
	// AVOptTypeImageSize wraps AV_OPT_TYPE_IMAGE_SIZE.
	AVOptTypeImageSize AVOptionType = C.AV_OPT_TYPE_IMAGE_SIZE
	// AVOptTypePixelFmt wraps AV_OPT_TYPE_PIXEL_FMT.
	AVOptTypePixelFmt AVOptionType = C.AV_OPT_TYPE_PIXEL_FMT
	// AVOptTypeSampleFmt wraps AV_OPT_TYPE_SAMPLE_FMT.
	AVOptTypeSampleFmt AVOptionType = C.AV_OPT_TYPE_SAMPLE_FMT
	// AVOptTypeVideoRate wraps AV_OPT_TYPE_VIDEO_RATE.
	AVOptTypeVideoRate AVOptionType = C.AV_OPT_TYPE_VIDEO_RATE
	// AVOptTypeDuration wraps AV_OPT_TYPE_DURATION.
	AVOptTypeDuration AVOptionType = C.AV_OPT_TYPE_DURATION
	// AVOptTypeColor wraps AV_OPT_TYPE_COLOR.
	AVOptTypeColor AVOptionType = C.AV_OPT_TYPE_COLOR
	// AVOptTypeChannelLayout wraps AV_OPT_TYPE_CHANNEL_LAYOUT.
	AVOptTypeChannelLayout AVOptionType = C.AV_OPT_TYPE_CHANNEL_LAYOUT
	// AVOptTypeBool wraps AV_OPT_TYPE_BOOL.
	AVOptTypeBool AVOptionType = C.AV_OPT_TYPE_BOOL
	// AVOptTypeChlayout wraps AV_OPT_TYPE_CHLAYOUT.
	AVOptTypeChlayout AVOptionType = C.AV_OPT_TYPE_CHLAYOUT
)

// --- Enum AVPixelFormat ---

// AVPixelFormat wraps AVPixelFormat.
type AVPixelFormat C.enum_AVPixelFormat

const (
	// AVPixFmtNone wraps AV_PIX_FMT_NONE.
	AVPixFmtNone AVPixelFormat = C.AV_PIX_FMT_NONE
	// AVPixFmtYuv420P wraps AV_PIX_FMT_YUV420P.
	AVPixFmtYuv420P AVPixelFormat = C.AV_PIX_FMT_YUV420P
	// AVPixFmtYuyv422 wraps AV_PIX_FMT_YUYV422.
	AVPixFmtYuyv422 AVPixelFormat = C.AV_PIX_FMT_YUYV422
	// AVPixFmtRgb24 wraps AV_PIX_FMT_RGB24.
	AVPixFmtRgb24 AVPixelFormat = C.AV_PIX_FMT_RGB24
	// AVPixFmtBgr24 wraps AV_PIX_FMT_BGR24.
	AVPixFmtBgr24 AVPixelFormat = C.AV_PIX_FMT_BGR24
	// AVPixFmtYuv422P wraps AV_PIX_FMT_YUV422P.
	AVPixFmtYuv422P AVPixelFormat = C.AV_PIX_FMT_YUV422P
	// AVPixFmtYuv444P wraps AV_PIX_FMT_YUV444P.
	AVPixFmtYuv444P AVPixelFormat = C.AV_PIX_FMT_YUV444P
	// AVPixFmtYuv410P wraps AV_PIX_FMT_YUV410P.
	AVPixFmtYuv410P AVPixelFormat = C.AV_PIX_FMT_YUV410P
	// AVPixFmtYuv411P wraps AV_PIX_FMT_YUV411P.
	AVPixFmtYuv411P AVPixelFormat = C.AV_PIX_FMT_YUV411P
	// AVPixFmtGray8 wraps AV_PIX_FMT_GRAY8.
	AVPixFmtGray8 AVPixelFormat = C.AV_PIX_FMT_GRAY8
	// AVPixFmtMonowhite wraps AV_PIX_FMT_MONOWHITE.
	AVPixFmtMonowhite AVPixelFormat = C.AV_PIX_FMT_MONOWHITE
	// AVPixFmtMonoblack wraps AV_PIX_FMT_MONOBLACK.
	AVPixFmtMonoblack AVPixelFormat = C.AV_PIX_FMT_MONOBLACK
	// AVPixFmtPal8 wraps AV_PIX_FMT_PAL8.
	AVPixFmtPal8 AVPixelFormat = C.AV_PIX_FMT_PAL8
	// AVPixFmtYuvj420P wraps AV_PIX_FMT_YUVJ420P.
	AVPixFmtYuvj420P AVPixelFormat = C.AV_PIX_FMT_YUVJ420P
	// AVPixFmtYuvj422P wraps AV_PIX_FMT_YUVJ422P.
	AVPixFmtYuvj422P AVPixelFormat = C.AV_PIX_FMT_YUVJ422P
	// AVPixFmtYuvj444P wraps AV_PIX_FMT_YUVJ444P.
	AVPixFmtYuvj444P AVPixelFormat = C.AV_PIX_FMT_YUVJ444P
	// AVPixFmtUyvy422 wraps AV_PIX_FMT_UYVY422.
	AVPixFmtUyvy422 AVPixelFormat = C.AV_PIX_FMT_UYVY422
	// AVPixFmtUyyvyy411 wraps AV_PIX_FMT_UYYVYY411.
	AVPixFmtUyyvyy411 AVPixelFormat = C.AV_PIX_FMT_UYYVYY411
	// AVPixFmtBgr8 wraps AV_PIX_FMT_BGR8.
	AVPixFmtBgr8 AVPixelFormat = C.AV_PIX_FMT_BGR8
	// AVPixFmtBgr4 wraps AV_PIX_FMT_BGR4.
	AVPixFmtBgr4 AVPixelFormat = C.AV_PIX_FMT_BGR4
	// AVPixFmtBgr4Byte wraps AV_PIX_FMT_BGR4_BYTE.
	AVPixFmtBgr4Byte AVPixelFormat = C.AV_PIX_FMT_BGR4_BYTE
	// AVPixFmtRgb8 wraps AV_PIX_FMT_RGB8.
	AVPixFmtRgb8 AVPixelFormat = C.AV_PIX_FMT_RGB8
	// AVPixFmtRgb4 wraps AV_PIX_FMT_RGB4.
	AVPixFmtRgb4 AVPixelFormat = C.AV_PIX_FMT_RGB4
	// AVPixFmtRgb4Byte wraps AV_PIX_FMT_RGB4_BYTE.
	AVPixFmtRgb4Byte AVPixelFormat = C.AV_PIX_FMT_RGB4_BYTE
	// AVPixFmtNv12 wraps AV_PIX_FMT_NV12.
	AVPixFmtNv12 AVPixelFormat = C.AV_PIX_FMT_NV12
	// AVPixFmtNv21 wraps AV_PIX_FMT_NV21.
	AVPixFmtNv21 AVPixelFormat = C.AV_PIX_FMT_NV21
	// AVPixFmtArgb wraps AV_PIX_FMT_ARGB.
	AVPixFmtArgb AVPixelFormat = C.AV_PIX_FMT_ARGB
	// AVPixFmtRgba wraps AV_PIX_FMT_RGBA.
	AVPixFmtRgba AVPixelFormat = C.AV_PIX_FMT_RGBA
	// AVPixFmtAbgr wraps AV_PIX_FMT_ABGR.
	AVPixFmtAbgr AVPixelFormat = C.AV_PIX_FMT_ABGR
	// AVPixFmtBgra wraps AV_PIX_FMT_BGRA.
	AVPixFmtBgra AVPixelFormat = C.AV_PIX_FMT_BGRA
	// AVPixFmtGray16Be wraps AV_PIX_FMT_GRAY16BE.
	AVPixFmtGray16Be AVPixelFormat = C.AV_PIX_FMT_GRAY16BE
	// AVPixFmtGray16Le wraps AV_PIX_FMT_GRAY16LE.
	AVPixFmtGray16Le AVPixelFormat = C.AV_PIX_FMT_GRAY16LE
	// AVPixFmtYuv440P wraps AV_PIX_FMT_YUV440P.
	AVPixFmtYuv440P AVPixelFormat = C.AV_PIX_FMT_YUV440P
	// AVPixFmtYuvj440P wraps AV_PIX_FMT_YUVJ440P.
	AVPixFmtYuvj440P AVPixelFormat = C.AV_PIX_FMT_YUVJ440P
	// AVPixFmtYuva420P wraps AV_PIX_FMT_YUVA420P.
	AVPixFmtYuva420P AVPixelFormat = C.AV_PIX_FMT_YUVA420P
	// AVPixFmtRgb48Be wraps AV_PIX_FMT_RGB48BE.
	AVPixFmtRgb48Be AVPixelFormat = C.AV_PIX_FMT_RGB48BE
	// AVPixFmtRgb48Le wraps AV_PIX_FMT_RGB48LE.
	AVPixFmtRgb48Le AVPixelFormat = C.AV_PIX_FMT_RGB48LE
	// AVPixFmtRgb565Be wraps AV_PIX_FMT_RGB565BE.
	AVPixFmtRgb565Be AVPixelFormat = C.AV_PIX_FMT_RGB565BE
	// AVPixFmtRgb565Le wraps AV_PIX_FMT_RGB565LE.
	AVPixFmtRgb565Le AVPixelFormat = C.AV_PIX_FMT_RGB565LE
	// AVPixFmtRgb555Be wraps AV_PIX_FMT_RGB555BE.
	AVPixFmtRgb555Be AVPixelFormat = C.AV_PIX_FMT_RGB555BE
	// AVPixFmtRgb555Le wraps AV_PIX_FMT_RGB555LE.
	AVPixFmtRgb555Le AVPixelFormat = C.AV_PIX_FMT_RGB555LE
	// AVPixFmtBgr565Be wraps AV_PIX_FMT_BGR565BE.
	AVPixFmtBgr565Be AVPixelFormat = C.AV_PIX_FMT_BGR565BE
	// AVPixFmtBgr565Le wraps AV_PIX_FMT_BGR565LE.
	AVPixFmtBgr565Le AVPixelFormat = C.AV_PIX_FMT_BGR565LE
	// AVPixFmtBgr555Be wraps AV_PIX_FMT_BGR555BE.
	AVPixFmtBgr555Be AVPixelFormat = C.AV_PIX_FMT_BGR555BE
	// AVPixFmtBgr555Le wraps AV_PIX_FMT_BGR555LE.
	AVPixFmtBgr555Le AVPixelFormat = C.AV_PIX_FMT_BGR555LE
	// AVPixFmtVaapi wraps AV_PIX_FMT_VAAPI.
	AVPixFmtVaapi AVPixelFormat = C.AV_PIX_FMT_VAAPI
	// AVPixFmtYuv420P16Le wraps AV_PIX_FMT_YUV420P16LE.
	AVPixFmtYuv420P16Le AVPixelFormat = C.AV_PIX_FMT_YUV420P16LE
	// AVPixFmtYuv420P16Be wraps AV_PIX_FMT_YUV420P16BE.
	AVPixFmtYuv420P16Be AVPixelFormat = C.AV_PIX_FMT_YUV420P16BE
	// AVPixFmtYuv422P16Le wraps AV_PIX_FMT_YUV422P16LE.
	AVPixFmtYuv422P16Le AVPixelFormat = C.AV_PIX_FMT_YUV422P16LE
	// AVPixFmtYuv422P16Be wraps AV_PIX_FMT_YUV422P16BE.
	AVPixFmtYuv422P16Be AVPixelFormat = C.AV_PIX_FMT_YUV422P16BE
	// AVPixFmtYuv444P16Le wraps AV_PIX_FMT_YUV444P16LE.
	AVPixFmtYuv444P16Le AVPixelFormat = C.AV_PIX_FMT_YUV444P16LE
	// AVPixFmtYuv444P16Be wraps AV_PIX_FMT_YUV444P16BE.
	AVPixFmtYuv444P16Be AVPixelFormat = C.AV_PIX_FMT_YUV444P16BE
	// AVPixFmtDxva2Vld wraps AV_PIX_FMT_DXVA2_VLD.
	AVPixFmtDxva2Vld AVPixelFormat = C.AV_PIX_FMT_DXVA2_VLD
	// AVPixFmtRgb444Le wraps AV_PIX_FMT_RGB444LE.
	AVPixFmtRgb444Le AVPixelFormat = C.AV_PIX_FMT_RGB444LE
	// AVPixFmtRgb444Be wraps AV_PIX_FMT_RGB444BE.
	AVPixFmtRgb444Be AVPixelFormat = C.AV_PIX_FMT_RGB444BE
	// AVPixFmtBgr444Le wraps AV_PIX_FMT_BGR444LE.
	AVPixFmtBgr444Le AVPixelFormat = C.AV_PIX_FMT_BGR444LE
	// AVPixFmtBgr444Be wraps AV_PIX_FMT_BGR444BE.
	AVPixFmtBgr444Be AVPixelFormat = C.AV_PIX_FMT_BGR444BE
	// AVPixFmtYa8 wraps AV_PIX_FMT_YA8.
	AVPixFmtYa8 AVPixelFormat = C.AV_PIX_FMT_YA8
	// AVPixFmtY400A wraps AV_PIX_FMT_Y400A.
	AVPixFmtY400A AVPixelFormat = C.AV_PIX_FMT_Y400A
	// AVPixFmtGray8A wraps AV_PIX_FMT_GRAY8A.
	AVPixFmtGray8A AVPixelFormat = C.AV_PIX_FMT_GRAY8A
	// AVPixFmtBgr48Be wraps AV_PIX_FMT_BGR48BE.
	AVPixFmtBgr48Be AVPixelFormat = C.AV_PIX_FMT_BGR48BE
	// AVPixFmtBgr48Le wraps AV_PIX_FMT_BGR48LE.
	AVPixFmtBgr48Le AVPixelFormat = C.AV_PIX_FMT_BGR48LE
	// AVPixFmtYuv420P9Be wraps AV_PIX_FMT_YUV420P9BE.
	AVPixFmtYuv420P9Be AVPixelFormat = C.AV_PIX_FMT_YUV420P9BE
	// AVPixFmtYuv420P9Le wraps AV_PIX_FMT_YUV420P9LE.
	AVPixFmtYuv420P9Le AVPixelFormat = C.AV_PIX_FMT_YUV420P9LE
	// AVPixFmtYuv420P10Be wraps AV_PIX_FMT_YUV420P10BE.
	AVPixFmtYuv420P10Be AVPixelFormat = C.AV_PIX_FMT_YUV420P10BE
	// AVPixFmtYuv420P10Le wraps AV_PIX_FMT_YUV420P10LE.
	AVPixFmtYuv420P10Le AVPixelFormat = C.AV_PIX_FMT_YUV420P10LE
	// AVPixFmtYuv422P10Be wraps AV_PIX_FMT_YUV422P10BE.
	AVPixFmtYuv422P10Be AVPixelFormat = C.AV_PIX_FMT_YUV422P10BE
	// AVPixFmtYuv422P10Le wraps AV_PIX_FMT_YUV422P10LE.
	AVPixFmtYuv422P10Le AVPixelFormat = C.AV_PIX_FMT_YUV422P10LE
	// AVPixFmtYuv444P9Be wraps AV_PIX_FMT_YUV444P9BE.
	AVPixFmtYuv444P9Be AVPixelFormat = C.AV_PIX_FMT_YUV444P9BE
	// AVPixFmtYuv444P9Le wraps AV_PIX_FMT_YUV444P9LE.
	AVPixFmtYuv444P9Le AVPixelFormat = C.AV_PIX_FMT_YUV444P9LE
	// AVPixFmtYuv444P10Be wraps AV_PIX_FMT_YUV444P10BE.
	AVPixFmtYuv444P10Be AVPixelFormat = C.AV_PIX_FMT_YUV444P10BE
	// AVPixFmtYuv444P10Le wraps AV_PIX_FMT_YUV444P10LE.
	AVPixFmtYuv444P10Le AVPixelFormat = C.AV_PIX_FMT_YUV444P10LE
	// AVPixFmtYuv422P9Be wraps AV_PIX_FMT_YUV422P9BE.
	AVPixFmtYuv422P9Be AVPixelFormat = C.AV_PIX_FMT_YUV422P9BE
	// AVPixFmtYuv422P9Le wraps AV_PIX_FMT_YUV422P9LE.
	AVPixFmtYuv422P9Le AVPixelFormat = C.AV_PIX_FMT_YUV422P9LE
	// AVPixFmtGbrp wraps AV_PIX_FMT_GBRP.
	AVPixFmtGbrp AVPixelFormat = C.AV_PIX_FMT_GBRP
	// AVPixFmtGbr24P wraps AV_PIX_FMT_GBR24P.
	AVPixFmtGbr24P AVPixelFormat = C.AV_PIX_FMT_GBR24P
	// AVPixFmtGbrp9Be wraps AV_PIX_FMT_GBRP9BE.
	AVPixFmtGbrp9Be AVPixelFormat = C.AV_PIX_FMT_GBRP9BE
	// AVPixFmtGbrp9Le wraps AV_PIX_FMT_GBRP9LE.
	AVPixFmtGbrp9Le AVPixelFormat = C.AV_PIX_FMT_GBRP9LE
	// AVPixFmtGbrp10Be wraps AV_PIX_FMT_GBRP10BE.
	AVPixFmtGbrp10Be AVPixelFormat = C.AV_PIX_FMT_GBRP10BE
	// AVPixFmtGbrp10Le wraps AV_PIX_FMT_GBRP10LE.
	AVPixFmtGbrp10Le AVPixelFormat = C.AV_PIX_FMT_GBRP10LE
	// AVPixFmtGbrp16Be wraps AV_PIX_FMT_GBRP16BE.
	AVPixFmtGbrp16Be AVPixelFormat = C.AV_PIX_FMT_GBRP16BE
	// AVPixFmtGbrp16Le wraps AV_PIX_FMT_GBRP16LE.
	AVPixFmtGbrp16Le AVPixelFormat = C.AV_PIX_FMT_GBRP16LE
	// AVPixFmtYuva422P wraps AV_PIX_FMT_YUVA422P.
	AVPixFmtYuva422P AVPixelFormat = C.AV_PIX_FMT_YUVA422P
	// AVPixFmtYuva444P wraps AV_PIX_FMT_YUVA444P.
	AVPixFmtYuva444P AVPixelFormat = C.AV_PIX_FMT_YUVA444P
	// AVPixFmtYuva420P9Be wraps AV_PIX_FMT_YUVA420P9BE.
	AVPixFmtYuva420P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P9BE
	// AVPixFmtYuva420P9Le wraps AV_PIX_FMT_YUVA420P9LE.
	AVPixFmtYuva420P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P9LE
	// AVPixFmtYuva422P9Be wraps AV_PIX_FMT_YUVA422P9BE.
	AVPixFmtYuva422P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P9BE
	// AVPixFmtYuva422P9Le wraps AV_PIX_FMT_YUVA422P9LE.
	AVPixFmtYuva422P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P9LE
	// AVPixFmtYuva444P9Be wraps AV_PIX_FMT_YUVA444P9BE.
	AVPixFmtYuva444P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P9BE
	// AVPixFmtYuva444P9Le wraps AV_PIX_FMT_YUVA444P9LE.
	AVPixFmtYuva444P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P9LE
	// AVPixFmtYuva420P10Be wraps AV_PIX_FMT_YUVA420P10BE.
	AVPixFmtYuva420P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P10BE
	// AVPixFmtYuva420P10Le wraps AV_PIX_FMT_YUVA420P10LE.
	AVPixFmtYuva420P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P10LE
	// AVPixFmtYuva422P10Be wraps AV_PIX_FMT_YUVA422P10BE.
	AVPixFmtYuva422P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P10BE
	// AVPixFmtYuva422P10Le wraps AV_PIX_FMT_YUVA422P10LE.
	AVPixFmtYuva422P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P10LE
	// AVPixFmtYuva444P10Be wraps AV_PIX_FMT_YUVA444P10BE.
	AVPixFmtYuva444P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P10BE
	// AVPixFmtYuva444P10Le wraps AV_PIX_FMT_YUVA444P10LE.
	AVPixFmtYuva444P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P10LE
	// AVPixFmtYuva420P16Be wraps AV_PIX_FMT_YUVA420P16BE.
	AVPixFmtYuva420P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P16BE
	// AVPixFmtYuva420P16Le wraps AV_PIX_FMT_YUVA420P16LE.
	AVPixFmtYuva420P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P16LE
	// AVPixFmtYuva422P16Be wraps AV_PIX_FMT_YUVA422P16BE.
	AVPixFmtYuva422P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P16BE
	// AVPixFmtYuva422P16Le wraps AV_PIX_FMT_YUVA422P16LE.
	AVPixFmtYuva422P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P16LE
	// AVPixFmtYuva444P16Be wraps AV_PIX_FMT_YUVA444P16BE.
	AVPixFmtYuva444P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P16BE
	// AVPixFmtYuva444P16Le wraps AV_PIX_FMT_YUVA444P16LE.
	AVPixFmtYuva444P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P16LE
	// AVPixFmtVdpau wraps AV_PIX_FMT_VDPAU.
	AVPixFmtVdpau AVPixelFormat = C.AV_PIX_FMT_VDPAU
	// AVPixFmtXyz12Le wraps AV_PIX_FMT_XYZ12LE.
	AVPixFmtXyz12Le AVPixelFormat = C.AV_PIX_FMT_XYZ12LE
	// AVPixFmtXyz12Be wraps AV_PIX_FMT_XYZ12BE.
	AVPixFmtXyz12Be AVPixelFormat = C.AV_PIX_FMT_XYZ12BE
	// AVPixFmtNv16 wraps AV_PIX_FMT_NV16.
	AVPixFmtNv16 AVPixelFormat = C.AV_PIX_FMT_NV16
	// AVPixFmtNv20Le wraps AV_PIX_FMT_NV20LE.
	AVPixFmtNv20Le AVPixelFormat = C.AV_PIX_FMT_NV20LE
	// AVPixFmtNv20Be wraps AV_PIX_FMT_NV20BE.
	AVPixFmtNv20Be AVPixelFormat = C.AV_PIX_FMT_NV20BE
	// AVPixFmtRgba64Be wraps AV_PIX_FMT_RGBA64BE.
	AVPixFmtRgba64Be AVPixelFormat = C.AV_PIX_FMT_RGBA64BE
	// AVPixFmtRgba64Le wraps AV_PIX_FMT_RGBA64LE.
	AVPixFmtRgba64Le AVPixelFormat = C.AV_PIX_FMT_RGBA64LE
	// AVPixFmtBgra64Be wraps AV_PIX_FMT_BGRA64BE.
	AVPixFmtBgra64Be AVPixelFormat = C.AV_PIX_FMT_BGRA64BE
	// AVPixFmtBgra64Le wraps AV_PIX_FMT_BGRA64LE.
	AVPixFmtBgra64Le AVPixelFormat = C.AV_PIX_FMT_BGRA64LE
	// AVPixFmtYvyu422 wraps AV_PIX_FMT_YVYU422.
	AVPixFmtYvyu422 AVPixelFormat = C.AV_PIX_FMT_YVYU422
	// AVPixFmtYa16Be wraps AV_PIX_FMT_YA16BE.
	AVPixFmtYa16Be AVPixelFormat = C.AV_PIX_FMT_YA16BE
	// AVPixFmtYa16Le wraps AV_PIX_FMT_YA16LE.
	AVPixFmtYa16Le AVPixelFormat = C.AV_PIX_FMT_YA16LE
	// AVPixFmtGbrap wraps AV_PIX_FMT_GBRAP.
	AVPixFmtGbrap AVPixelFormat = C.AV_PIX_FMT_GBRAP
	// AVPixFmtGbrap16Be wraps AV_PIX_FMT_GBRAP16BE.
	AVPixFmtGbrap16Be AVPixelFormat = C.AV_PIX_FMT_GBRAP16BE
	// AVPixFmtGbrap16Le wraps AV_PIX_FMT_GBRAP16LE.
	AVPixFmtGbrap16Le AVPixelFormat = C.AV_PIX_FMT_GBRAP16LE
	// AVPixFmtQsv wraps AV_PIX_FMT_QSV.
	AVPixFmtQsv AVPixelFormat = C.AV_PIX_FMT_QSV
	// AVPixFmtMmal wraps AV_PIX_FMT_MMAL.
	AVPixFmtMmal AVPixelFormat = C.AV_PIX_FMT_MMAL
	// AVPixFmtD3D11VaVld wraps AV_PIX_FMT_D3D11VA_VLD.
	AVPixFmtD3D11VaVld AVPixelFormat = C.AV_PIX_FMT_D3D11VA_VLD
	// AVPixFmtCuda wraps AV_PIX_FMT_CUDA.
	AVPixFmtCuda AVPixelFormat = C.AV_PIX_FMT_CUDA
	// AVPixFmt0Rgb wraps AV_PIX_FMT_0RGB.
	AVPixFmt0Rgb AVPixelFormat = C.AV_PIX_FMT_0RGB
	// AVPixFmtRgb0 wraps AV_PIX_FMT_RGB0.
	AVPixFmtRgb0 AVPixelFormat = C.AV_PIX_FMT_RGB0
	// AVPixFmt0Bgr wraps AV_PIX_FMT_0BGR.
	AVPixFmt0Bgr AVPixelFormat = C.AV_PIX_FMT_0BGR
	// AVPixFmtBgr0 wraps AV_PIX_FMT_BGR0.
	AVPixFmtBgr0 AVPixelFormat = C.AV_PIX_FMT_BGR0
	// AVPixFmtYuv420P12Be wraps AV_PIX_FMT_YUV420P12BE.
	AVPixFmtYuv420P12Be AVPixelFormat = C.AV_PIX_FMT_YUV420P12BE
	// AVPixFmtYuv420P12Le wraps AV_PIX_FMT_YUV420P12LE.
	AVPixFmtYuv420P12Le AVPixelFormat = C.AV_PIX_FMT_YUV420P12LE
	// AVPixFmtYuv420P14Be wraps AV_PIX_FMT_YUV420P14BE.
	AVPixFmtYuv420P14Be AVPixelFormat = C.AV_PIX_FMT_YUV420P14BE
	// AVPixFmtYuv420P14Le wraps AV_PIX_FMT_YUV420P14LE.
	AVPixFmtYuv420P14Le AVPixelFormat = C.AV_PIX_FMT_YUV420P14LE
	// AVPixFmtYuv422P12Be wraps AV_PIX_FMT_YUV422P12BE.
	AVPixFmtYuv422P12Be AVPixelFormat = C.AV_PIX_FMT_YUV422P12BE
	// AVPixFmtYuv422P12Le wraps AV_PIX_FMT_YUV422P12LE.
	AVPixFmtYuv422P12Le AVPixelFormat = C.AV_PIX_FMT_YUV422P12LE
	// AVPixFmtYuv422P14Be wraps AV_PIX_FMT_YUV422P14BE.
	AVPixFmtYuv422P14Be AVPixelFormat = C.AV_PIX_FMT_YUV422P14BE
	// AVPixFmtYuv422P14Le wraps AV_PIX_FMT_YUV422P14LE.
	AVPixFmtYuv422P14Le AVPixelFormat = C.AV_PIX_FMT_YUV422P14LE
	// AVPixFmtYuv444P12Be wraps AV_PIX_FMT_YUV444P12BE.
	AVPixFmtYuv444P12Be AVPixelFormat = C.AV_PIX_FMT_YUV444P12BE
	// AVPixFmtYuv444P12Le wraps AV_PIX_FMT_YUV444P12LE.
	AVPixFmtYuv444P12Le AVPixelFormat = C.AV_PIX_FMT_YUV444P12LE
	// AVPixFmtYuv444P14Be wraps AV_PIX_FMT_YUV444P14BE.
	AVPixFmtYuv444P14Be AVPixelFormat = C.AV_PIX_FMT_YUV444P14BE
	// AVPixFmtYuv444P14Le wraps AV_PIX_FMT_YUV444P14LE.
	AVPixFmtYuv444P14Le AVPixelFormat = C.AV_PIX_FMT_YUV444P14LE
	// AVPixFmtGbrp12Be wraps AV_PIX_FMT_GBRP12BE.
	AVPixFmtGbrp12Be AVPixelFormat = C.AV_PIX_FMT_GBRP12BE
	// AVPixFmtGbrp12Le wraps AV_PIX_FMT_GBRP12LE.
	AVPixFmtGbrp12Le AVPixelFormat = C.AV_PIX_FMT_GBRP12LE
	// AVPixFmtGbrp14Be wraps AV_PIX_FMT_GBRP14BE.
	AVPixFmtGbrp14Be AVPixelFormat = C.AV_PIX_FMT_GBRP14BE
	// AVPixFmtGbrp14Le wraps AV_PIX_FMT_GBRP14LE.
	AVPixFmtGbrp14Le AVPixelFormat = C.AV_PIX_FMT_GBRP14LE
	// AVPixFmtYuvj411P wraps AV_PIX_FMT_YUVJ411P.
	AVPixFmtYuvj411P AVPixelFormat = C.AV_PIX_FMT_YUVJ411P
	// AVPixFmtBayerBggr8 wraps AV_PIX_FMT_BAYER_BGGR8.
	AVPixFmtBayerBggr8 AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR8
	// AVPixFmtBayerRggb8 wraps AV_PIX_FMT_BAYER_RGGB8.
	AVPixFmtBayerRggb8 AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB8
	// AVPixFmtBayerGbrg8 wraps AV_PIX_FMT_BAYER_GBRG8.
	AVPixFmtBayerGbrg8 AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG8
	// AVPixFmtBayerGrbg8 wraps AV_PIX_FMT_BAYER_GRBG8.
	AVPixFmtBayerGrbg8 AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG8
	// AVPixFmtBayerBggr16Le wraps AV_PIX_FMT_BAYER_BGGR16LE.
	AVPixFmtBayerBggr16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16LE
	// AVPixFmtBayerBggr16Be wraps AV_PIX_FMT_BAYER_BGGR16BE.
	AVPixFmtBayerBggr16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16BE
	// AVPixFmtBayerRggb16Le wraps AV_PIX_FMT_BAYER_RGGB16LE.
	AVPixFmtBayerRggb16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16LE
	// AVPixFmtBayerRggb16Be wraps AV_PIX_FMT_BAYER_RGGB16BE.
	AVPixFmtBayerRggb16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16BE
	// AVPixFmtBayerGbrg16Le wraps AV_PIX_FMT_BAYER_GBRG16LE.
	AVPixFmtBayerGbrg16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16LE
	// AVPixFmtBayerGbrg16Be wraps AV_PIX_FMT_BAYER_GBRG16BE.
	AVPixFmtBayerGbrg16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16BE
	// AVPixFmtBayerGrbg16Le wraps AV_PIX_FMT_BAYER_GRBG16LE.
	AVPixFmtBayerGrbg16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16LE
	// AVPixFmtBayerGrbg16Be wraps AV_PIX_FMT_BAYER_GRBG16BE.
	AVPixFmtBayerGrbg16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16BE
	// AVPixFmtXvmc wraps AV_PIX_FMT_XVMC.
	AVPixFmtXvmc AVPixelFormat = C.AV_PIX_FMT_XVMC
	// AVPixFmtYuv440P10Le wraps AV_PIX_FMT_YUV440P10LE.
	AVPixFmtYuv440P10Le AVPixelFormat = C.AV_PIX_FMT_YUV440P10LE
	// AVPixFmtYuv440P10Be wraps AV_PIX_FMT_YUV440P10BE.
	AVPixFmtYuv440P10Be AVPixelFormat = C.AV_PIX_FMT_YUV440P10BE
	// AVPixFmtYuv440P12Le wraps AV_PIX_FMT_YUV440P12LE.
	AVPixFmtYuv440P12Le AVPixelFormat = C.AV_PIX_FMT_YUV440P12LE
	// AVPixFmtYuv440P12Be wraps AV_PIX_FMT_YUV440P12BE.
	AVPixFmtYuv440P12Be AVPixelFormat = C.AV_PIX_FMT_YUV440P12BE
	// AVPixFmtAyuv64Le wraps AV_PIX_FMT_AYUV64LE.
	AVPixFmtAyuv64Le AVPixelFormat = C.AV_PIX_FMT_AYUV64LE
	// AVPixFmtAyuv64Be wraps AV_PIX_FMT_AYUV64BE.
	AVPixFmtAyuv64Be AVPixelFormat = C.AV_PIX_FMT_AYUV64BE
	// AVPixFmtVideotoolbox wraps AV_PIX_FMT_VIDEOTOOLBOX.
	AVPixFmtVideotoolbox AVPixelFormat = C.AV_PIX_FMT_VIDEOTOOLBOX
	// AVPixFmtP010Le wraps AV_PIX_FMT_P010LE.
	AVPixFmtP010Le AVPixelFormat = C.AV_PIX_FMT_P010LE
	// AVPixFmtP010Be wraps AV_PIX_FMT_P010BE.
	AVPixFmtP010Be AVPixelFormat = C.AV_PIX_FMT_P010BE
	// AVPixFmtGbrap12Be wraps AV_PIX_FMT_GBRAP12BE.
	AVPixFmtGbrap12Be AVPixelFormat = C.AV_PIX_FMT_GBRAP12BE
	// AVPixFmtGbrap12Le wraps AV_PIX_FMT_GBRAP12LE.
	AVPixFmtGbrap12Le AVPixelFormat = C.AV_PIX_FMT_GBRAP12LE
	// AVPixFmtGbrap10Be wraps AV_PIX_FMT_GBRAP10BE.
	AVPixFmtGbrap10Be AVPixelFormat = C.AV_PIX_FMT_GBRAP10BE
	// AVPixFmtGbrap10Le wraps AV_PIX_FMT_GBRAP10LE.
	AVPixFmtGbrap10Le AVPixelFormat = C.AV_PIX_FMT_GBRAP10LE
	// AVPixFmtMediacodec wraps AV_PIX_FMT_MEDIACODEC.
	AVPixFmtMediacodec AVPixelFormat = C.AV_PIX_FMT_MEDIACODEC
	// AVPixFmtGray12Be wraps AV_PIX_FMT_GRAY12BE.
	AVPixFmtGray12Be AVPixelFormat = C.AV_PIX_FMT_GRAY12BE
	// AVPixFmtGray12Le wraps AV_PIX_FMT_GRAY12LE.
	AVPixFmtGray12Le AVPixelFormat = C.AV_PIX_FMT_GRAY12LE
	// AVPixFmtGray10Be wraps AV_PIX_FMT_GRAY10BE.
	AVPixFmtGray10Be AVPixelFormat = C.AV_PIX_FMT_GRAY10BE
	// AVPixFmtGray10Le wraps AV_PIX_FMT_GRAY10LE.
	AVPixFmtGray10Le AVPixelFormat = C.AV_PIX_FMT_GRAY10LE
	// AVPixFmtP016Le wraps AV_PIX_FMT_P016LE.
	AVPixFmtP016Le AVPixelFormat = C.AV_PIX_FMT_P016LE
	// AVPixFmtP016Be wraps AV_PIX_FMT_P016BE.
	AVPixFmtP016Be AVPixelFormat = C.AV_PIX_FMT_P016BE
	// AVPixFmtD3D11 wraps AV_PIX_FMT_D3D11.
	AVPixFmtD3D11 AVPixelFormat = C.AV_PIX_FMT_D3D11
	// AVPixFmtGray9Be wraps AV_PIX_FMT_GRAY9BE.
	AVPixFmtGray9Be AVPixelFormat = C.AV_PIX_FMT_GRAY9BE
	// AVPixFmtGray9Le wraps AV_PIX_FMT_GRAY9LE.
	AVPixFmtGray9Le AVPixelFormat = C.AV_PIX_FMT_GRAY9LE
	// AVPixFmtGbrpf32Be wraps AV_PIX_FMT_GBRPF32BE.
	AVPixFmtGbrpf32Be AVPixelFormat = C.AV_PIX_FMT_GBRPF32BE
	// AVPixFmtGbrpf32Le wraps AV_PIX_FMT_GBRPF32LE.
	AVPixFmtGbrpf32Le AVPixelFormat = C.AV_PIX_FMT_GBRPF32LE
	// AVPixFmtGbrapf32Be wraps AV_PIX_FMT_GBRAPF32BE.
	AVPixFmtGbrapf32Be AVPixelFormat = C.AV_PIX_FMT_GBRAPF32BE
	// AVPixFmtGbrapf32Le wraps AV_PIX_FMT_GBRAPF32LE.
	AVPixFmtGbrapf32Le AVPixelFormat = C.AV_PIX_FMT_GBRAPF32LE
	// AVPixFmtDrmPrime wraps AV_PIX_FMT_DRM_PRIME.
	AVPixFmtDrmPrime AVPixelFormat = C.AV_PIX_FMT_DRM_PRIME
	// AVPixFmtOpencl wraps AV_PIX_FMT_OPENCL.
	AVPixFmtOpencl AVPixelFormat = C.AV_PIX_FMT_OPENCL
	// AVPixFmtGray14Be wraps AV_PIX_FMT_GRAY14BE.
	AVPixFmtGray14Be AVPixelFormat = C.AV_PIX_FMT_GRAY14BE
	// AVPixFmtGray14Le wraps AV_PIX_FMT_GRAY14LE.
	AVPixFmtGray14Le AVPixelFormat = C.AV_PIX_FMT_GRAY14LE
	// AVPixFmtGrayf32Be wraps AV_PIX_FMT_GRAYF32BE.
	AVPixFmtGrayf32Be AVPixelFormat = C.AV_PIX_FMT_GRAYF32BE
	// AVPixFmtGrayf32Le wraps AV_PIX_FMT_GRAYF32LE.
	AVPixFmtGrayf32Le AVPixelFormat = C.AV_PIX_FMT_GRAYF32LE
	// AVPixFmtYuva422P12Be wraps AV_PIX_FMT_YUVA422P12BE.
	AVPixFmtYuva422P12Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P12BE
	// AVPixFmtYuva422P12Le wraps AV_PIX_FMT_YUVA422P12LE.
	AVPixFmtYuva422P12Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P12LE
	// AVPixFmtYuva444P12Be wraps AV_PIX_FMT_YUVA444P12BE.
	AVPixFmtYuva444P12Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P12BE
	// AVPixFmtYuva444P12Le wraps AV_PIX_FMT_YUVA444P12LE.
	AVPixFmtYuva444P12Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P12LE
	// AVPixFmtNv24 wraps AV_PIX_FMT_NV24.
	AVPixFmtNv24 AVPixelFormat = C.AV_PIX_FMT_NV24
	// AVPixFmtNv42 wraps AV_PIX_FMT_NV42.
	AVPixFmtNv42 AVPixelFormat = C.AV_PIX_FMT_NV42
	// AVPixFmtVulkan wraps AV_PIX_FMT_VULKAN.
	AVPixFmtVulkan AVPixelFormat = C.AV_PIX_FMT_VULKAN
	// AVPixFmtY210Be wraps AV_PIX_FMT_Y210BE.
	AVPixFmtY210Be AVPixelFormat = C.AV_PIX_FMT_Y210BE
	// AVPixFmtY210Le wraps AV_PIX_FMT_Y210LE.
	AVPixFmtY210Le AVPixelFormat = C.AV_PIX_FMT_Y210LE
	// AVPixFmtX2Rgb10Le wraps AV_PIX_FMT_X2RGB10LE.
	AVPixFmtX2Rgb10Le AVPixelFormat = C.AV_PIX_FMT_X2RGB10LE
	// AVPixFmtX2Rgb10Be wraps AV_PIX_FMT_X2RGB10BE.
	AVPixFmtX2Rgb10Be AVPixelFormat = C.AV_PIX_FMT_X2RGB10BE
	// AVPixFmtX2Bgr10Le wraps AV_PIX_FMT_X2BGR10LE.
	AVPixFmtX2Bgr10Le AVPixelFormat = C.AV_PIX_FMT_X2BGR10LE
	// AVPixFmtX2Bgr10Be wraps AV_PIX_FMT_X2BGR10BE.
	AVPixFmtX2Bgr10Be AVPixelFormat = C.AV_PIX_FMT_X2BGR10BE
	// AVPixFmtP210Be wraps AV_PIX_FMT_P210BE.
	AVPixFmtP210Be AVPixelFormat = C.AV_PIX_FMT_P210BE
	// AVPixFmtP210Le wraps AV_PIX_FMT_P210LE.
	AVPixFmtP210Le AVPixelFormat = C.AV_PIX_FMT_P210LE
	// AVPixFmtP410Be wraps AV_PIX_FMT_P410BE.
	AVPixFmtP410Be AVPixelFormat = C.AV_PIX_FMT_P410BE
	// AVPixFmtP410Le wraps AV_PIX_FMT_P410LE.
	AVPixFmtP410Le AVPixelFormat = C.AV_PIX_FMT_P410LE
	// AVPixFmtP216Be wraps AV_PIX_FMT_P216BE.
	AVPixFmtP216Be AVPixelFormat = C.AV_PIX_FMT_P216BE
	// AVPixFmtP216Le wraps AV_PIX_FMT_P216LE.
	AVPixFmtP216Le AVPixelFormat = C.AV_PIX_FMT_P216LE
	// AVPixFmtP416Be wraps AV_PIX_FMT_P416BE.
	AVPixFmtP416Be AVPixelFormat = C.AV_PIX_FMT_P416BE
	// AVPixFmtP416Le wraps AV_PIX_FMT_P416LE.
	AVPixFmtP416Le AVPixelFormat = C.AV_PIX_FMT_P416LE
	// AVPixFmtVuya wraps AV_PIX_FMT_VUYA.
	AVPixFmtVuya AVPixelFormat = C.AV_PIX_FMT_VUYA
	// AVPixFmtRgbaf16Be wraps AV_PIX_FMT_RGBAF16BE.
	AVPixFmtRgbaf16Be AVPixelFormat = C.AV_PIX_FMT_RGBAF16BE
	// AVPixFmtRgbaf16Le wraps AV_PIX_FMT_RGBAF16LE.
	AVPixFmtRgbaf16Le AVPixelFormat = C.AV_PIX_FMT_RGBAF16LE
	// AVPixFmtVuyx wraps AV_PIX_FMT_VUYX.
	AVPixFmtVuyx AVPixelFormat = C.AV_PIX_FMT_VUYX
	// AVPixFmtP012Le wraps AV_PIX_FMT_P012LE.
	AVPixFmtP012Le AVPixelFormat = C.AV_PIX_FMT_P012LE
	// AVPixFmtP012Be wraps AV_PIX_FMT_P012BE.
	AVPixFmtP012Be AVPixelFormat = C.AV_PIX_FMT_P012BE
	// AVPixFmtY212Be wraps AV_PIX_FMT_Y212BE.
	AVPixFmtY212Be AVPixelFormat = C.AV_PIX_FMT_Y212BE
	// AVPixFmtY212Le wraps AV_PIX_FMT_Y212LE.
	AVPixFmtY212Le AVPixelFormat = C.AV_PIX_FMT_Y212LE
	// AVPixFmtXv30Be wraps AV_PIX_FMT_XV30BE.
	AVPixFmtXv30Be AVPixelFormat = C.AV_PIX_FMT_XV30BE
	// AVPixFmtXv30Le wraps AV_PIX_FMT_XV30LE.
	AVPixFmtXv30Le AVPixelFormat = C.AV_PIX_FMT_XV30LE
	// AVPixFmtXv36Be wraps AV_PIX_FMT_XV36BE.
	AVPixFmtXv36Be AVPixelFormat = C.AV_PIX_FMT_XV36BE
	// AVPixFmtXv36Le wraps AV_PIX_FMT_XV36LE.
	AVPixFmtXv36Le AVPixelFormat = C.AV_PIX_FMT_XV36LE
	// AVPixFmtRgbf32Be wraps AV_PIX_FMT_RGBF32BE.
	AVPixFmtRgbf32Be AVPixelFormat = C.AV_PIX_FMT_RGBF32BE
	// AVPixFmtRgbf32Le wraps AV_PIX_FMT_RGBF32LE.
	AVPixFmtRgbf32Le AVPixelFormat = C.AV_PIX_FMT_RGBF32LE
	// AVPixFmtRgbaf32Be wraps AV_PIX_FMT_RGBAF32BE.
	AVPixFmtRgbaf32Be AVPixelFormat = C.AV_PIX_FMT_RGBAF32BE
	// AVPixFmtRgbaf32Le wraps AV_PIX_FMT_RGBAF32LE.
	AVPixFmtRgbaf32Le AVPixelFormat = C.AV_PIX_FMT_RGBAF32LE
	// AVPixFmtNb wraps AV_PIX_FMT_NB.
	AVPixFmtNb AVPixelFormat = C.AV_PIX_FMT_NB
)

// --- Enum AVColorPrimaries ---

// AVColorPrimaries wraps AVColorPrimaries.
type AVColorPrimaries C.enum_AVColorPrimaries

const (
	// AVColPriReserved0 wraps AVCOL_PRI_RESERVED0.
	AVColPriReserved0 AVColorPrimaries = C.AVCOL_PRI_RESERVED0
	// AVColPriBt709 wraps AVCOL_PRI_BT709.
	AVColPriBt709 AVColorPrimaries = C.AVCOL_PRI_BT709
	// AVColPriUnspecified wraps AVCOL_PRI_UNSPECIFIED.
	AVColPriUnspecified AVColorPrimaries = C.AVCOL_PRI_UNSPECIFIED
	// AVColPriReserved wraps AVCOL_PRI_RESERVED.
	AVColPriReserved AVColorPrimaries = C.AVCOL_PRI_RESERVED
	// AVColPriBt470M wraps AVCOL_PRI_BT470M.
	AVColPriBt470M AVColorPrimaries = C.AVCOL_PRI_BT470M
	// AVColPriBt470Bg wraps AVCOL_PRI_BT470BG.
	AVColPriBt470Bg AVColorPrimaries = C.AVCOL_PRI_BT470BG
	// AVColPriSmpte170M wraps AVCOL_PRI_SMPTE170M.
	AVColPriSmpte170M AVColorPrimaries = C.AVCOL_PRI_SMPTE170M
	// AVColPriSmpte240M wraps AVCOL_PRI_SMPTE240M.
	AVColPriSmpte240M AVColorPrimaries = C.AVCOL_PRI_SMPTE240M
	// AVColPriFilm wraps AVCOL_PRI_FILM.
	AVColPriFilm AVColorPrimaries = C.AVCOL_PRI_FILM
	// AVColPriBt2020 wraps AVCOL_PRI_BT2020.
	AVColPriBt2020 AVColorPrimaries = C.AVCOL_PRI_BT2020
	// AVColPriSmpte428 wraps AVCOL_PRI_SMPTE428.
	AVColPriSmpte428 AVColorPrimaries = C.AVCOL_PRI_SMPTE428
	// AVColPriSmptest4281 wraps AVCOL_PRI_SMPTEST428_1.
	AVColPriSmptest4281 AVColorPrimaries = C.AVCOL_PRI_SMPTEST428_1
	// AVColPriSmpte431 wraps AVCOL_PRI_SMPTE431.
	AVColPriSmpte431 AVColorPrimaries = C.AVCOL_PRI_SMPTE431
	// AVColPriSmpte432 wraps AVCOL_PRI_SMPTE432.
	AVColPriSmpte432 AVColorPrimaries = C.AVCOL_PRI_SMPTE432
	// AVColPriEbu3213 wraps AVCOL_PRI_EBU3213.
	AVColPriEbu3213 AVColorPrimaries = C.AVCOL_PRI_EBU3213
	// AVColPriJedecP22 wraps AVCOL_PRI_JEDEC_P22.
	AVColPriJedecP22 AVColorPrimaries = C.AVCOL_PRI_JEDEC_P22
	// AVColPriNb wraps AVCOL_PRI_NB.
	AVColPriNb AVColorPrimaries = C.AVCOL_PRI_NB
)

// --- Enum AVColorTransferCharacteristic ---

// AVColorTransferCharacteristic wraps AVColorTransferCharacteristic.
type AVColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

const (
	// AVColTrcReserved0 wraps AVCOL_TRC_RESERVED0.
	AVColTrcReserved0 AVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED0
	// AVColTrcBt709 wraps AVCOL_TRC_BT709.
	AVColTrcBt709 AVColorTransferCharacteristic = C.AVCOL_TRC_BT709
	// AVColTrcUnspecified wraps AVCOL_TRC_UNSPECIFIED.
	AVColTrcUnspecified AVColorTransferCharacteristic = C.AVCOL_TRC_UNSPECIFIED
	// AVColTrcReserved wraps AVCOL_TRC_RESERVED.
	AVColTrcReserved AVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED
	// AVColTrcGamma22 wraps AVCOL_TRC_GAMMA22.
	AVColTrcGamma22 AVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA22
	// AVColTrcGamma28 wraps AVCOL_TRC_GAMMA28.
	AVColTrcGamma28 AVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA28
	// AVColTrcSmpte170M wraps AVCOL_TRC_SMPTE170M.
	AVColTrcSmpte170M AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE170M
	// AVColTrcSmpte240M wraps AVCOL_TRC_SMPTE240M.
	AVColTrcSmpte240M AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE240M
	// AVColTrcLinear wraps AVCOL_TRC_LINEAR.
	AVColTrcLinear AVColorTransferCharacteristic = C.AVCOL_TRC_LINEAR
	// AVColTrcLog wraps AVCOL_TRC_LOG.
	AVColTrcLog AVColorTransferCharacteristic = C.AVCOL_TRC_LOG
	// AVColTrcLogSqrt wraps AVCOL_TRC_LOG_SQRT.
	AVColTrcLogSqrt AVColorTransferCharacteristic = C.AVCOL_TRC_LOG_SQRT
	// AVColTrcIec6196624 wraps AVCOL_TRC_IEC61966_2_4.
	AVColTrcIec6196624 AVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_4
	// AVColTrcBt1361Ecg wraps AVCOL_TRC_BT1361_ECG.
	AVColTrcBt1361Ecg AVColorTransferCharacteristic = C.AVCOL_TRC_BT1361_ECG
	// AVColTrcIec6196621 wraps AVCOL_TRC_IEC61966_2_1.
	AVColTrcIec6196621 AVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_1
	// AVColTrcBt202010 wraps AVCOL_TRC_BT2020_10.
	AVColTrcBt202010 AVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_10
	// AVColTrcBt202012 wraps AVCOL_TRC_BT2020_12.
	AVColTrcBt202012 AVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_12
	// AVColTrcSmpte2084 wraps AVCOL_TRC_SMPTE2084.
	AVColTrcSmpte2084 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE2084
	// AVColTrcSmptest2084 wraps AVCOL_TRC_SMPTEST2084.
	AVColTrcSmptest2084 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST2084
	// AVColTrcSmpte428 wraps AVCOL_TRC_SMPTE428.
	AVColTrcSmpte428 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE428
	// AVColTrcSmptest4281 wraps AVCOL_TRC_SMPTEST428_1.
	AVColTrcSmptest4281 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST428_1
	// AVColTrcAribStdB67 wraps AVCOL_TRC_ARIB_STD_B67.
	AVColTrcAribStdB67 AVColorTransferCharacteristic = C.AVCOL_TRC_ARIB_STD_B67
	// AVColTrcNb wraps AVCOL_TRC_NB.
	AVColTrcNb AVColorTransferCharacteristic = C.AVCOL_TRC_NB
)

// --- Enum AVColorSpace ---

// AVColorSpace wraps AVColorSpace.
type AVColorSpace C.enum_AVColorSpace

const (
	// AVColSpcRgb wraps AVCOL_SPC_RGB.
	AVColSpcRgb AVColorSpace = C.AVCOL_SPC_RGB
	// AVColSpcBt709 wraps AVCOL_SPC_BT709.
	AVColSpcBt709 AVColorSpace = C.AVCOL_SPC_BT709
	// AVColSpcUnspecified wraps AVCOL_SPC_UNSPECIFIED.
	AVColSpcUnspecified AVColorSpace = C.AVCOL_SPC_UNSPECIFIED
	// AVColSpcReserved wraps AVCOL_SPC_RESERVED.
	AVColSpcReserved AVColorSpace = C.AVCOL_SPC_RESERVED
	// AVColSpcFcc wraps AVCOL_SPC_FCC.
	AVColSpcFcc AVColorSpace = C.AVCOL_SPC_FCC
	// AVColSpcBt470Bg wraps AVCOL_SPC_BT470BG.
	AVColSpcBt470Bg AVColorSpace = C.AVCOL_SPC_BT470BG
	// AVColSpcSmpte170M wraps AVCOL_SPC_SMPTE170M.
	AVColSpcSmpte170M AVColorSpace = C.AVCOL_SPC_SMPTE170M
	// AVColSpcSmpte240M wraps AVCOL_SPC_SMPTE240M.
	AVColSpcSmpte240M AVColorSpace = C.AVCOL_SPC_SMPTE240M
	// AVColSpcYcgco wraps AVCOL_SPC_YCGCO.
	AVColSpcYcgco AVColorSpace = C.AVCOL_SPC_YCGCO
	// AVColSpcYcocg wraps AVCOL_SPC_YCOCG.
	AVColSpcYcocg AVColorSpace = C.AVCOL_SPC_YCOCG
	// AVColSpcBt2020Ncl wraps AVCOL_SPC_BT2020_NCL.
	AVColSpcBt2020Ncl AVColorSpace = C.AVCOL_SPC_BT2020_NCL
	// AVColSpcBt2020Cl wraps AVCOL_SPC_BT2020_CL.
	AVColSpcBt2020Cl AVColorSpace = C.AVCOL_SPC_BT2020_CL
	// AVColSpcSmpte2085 wraps AVCOL_SPC_SMPTE2085.
	AVColSpcSmpte2085 AVColorSpace = C.AVCOL_SPC_SMPTE2085
	// AVColSpcChromaDerivedNcl wraps AVCOL_SPC_CHROMA_DERIVED_NCL.
	AVColSpcChromaDerivedNcl AVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_NCL
	// AVColSpcChromaDerivedCl wraps AVCOL_SPC_CHROMA_DERIVED_CL.
	AVColSpcChromaDerivedCl AVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_CL
	// AVColSpcIctcp wraps AVCOL_SPC_ICTCP.
	AVColSpcIctcp AVColorSpace = C.AVCOL_SPC_ICTCP
	// AVColSpcNb wraps AVCOL_SPC_NB.
	AVColSpcNb AVColorSpace = C.AVCOL_SPC_NB
)

// --- Enum AVColorRange ---

// AVColorRange wraps AVColorRange.
type AVColorRange C.enum_AVColorRange

const (
	// AVColRangeUnspecified wraps AVCOL_RANGE_UNSPECIFIED.
	AVColRangeUnspecified AVColorRange = C.AVCOL_RANGE_UNSPECIFIED
	// AVColRangeMpeg wraps AVCOL_RANGE_MPEG.
	AVColRangeMpeg AVColorRange = C.AVCOL_RANGE_MPEG
	// AVColRangeJpeg wraps AVCOL_RANGE_JPEG.
	AVColRangeJpeg AVColorRange = C.AVCOL_RANGE_JPEG
	// AVColRangeNb wraps AVCOL_RANGE_NB.
	AVColRangeNb AVColorRange = C.AVCOL_RANGE_NB
)

// --- Enum AVChromaLocation ---

// AVChromaLocation wraps AVChromaLocation.
type AVChromaLocation C.enum_AVChromaLocation

const (
	// AVChromaLocUnspecified wraps AVCHROMA_LOC_UNSPECIFIED.
	AVChromaLocUnspecified AVChromaLocation = C.AVCHROMA_LOC_UNSPECIFIED
	// AVChromaLocLeft wraps AVCHROMA_LOC_LEFT.
	AVChromaLocLeft AVChromaLocation = C.AVCHROMA_LOC_LEFT
	// AVChromaLocCenter wraps AVCHROMA_LOC_CENTER.
	AVChromaLocCenter AVChromaLocation = C.AVCHROMA_LOC_CENTER
	// AVChromaLocTopleft wraps AVCHROMA_LOC_TOPLEFT.
	AVChromaLocTopleft AVChromaLocation = C.AVCHROMA_LOC_TOPLEFT
	// AVChromaLocTop wraps AVCHROMA_LOC_TOP.
	AVChromaLocTop AVChromaLocation = C.AVCHROMA_LOC_TOP
	// AVChromaLocBottomleft wraps AVCHROMA_LOC_BOTTOMLEFT.
	AVChromaLocBottomleft AVChromaLocation = C.AVCHROMA_LOC_BOTTOMLEFT
	// AVChromaLocBottom wraps AVCHROMA_LOC_BOTTOM.
	AVChromaLocBottom AVChromaLocation = C.AVCHROMA_LOC_BOTTOM
	// AVChromaLocNb wraps AVCHROMA_LOC_NB.
	AVChromaLocNb AVChromaLocation = C.AVCHROMA_LOC_NB
)

// --- Enum AVSampleFormat ---

// AVSampleFormat wraps AVSampleFormat.
type AVSampleFormat C.enum_AVSampleFormat

const (
	// AVSampleFmtNone wraps AV_SAMPLE_FMT_NONE.
	AVSampleFmtNone AVSampleFormat = C.AV_SAMPLE_FMT_NONE
	// AVSampleFmtU8 wraps AV_SAMPLE_FMT_U8.
	AVSampleFmtU8 AVSampleFormat = C.AV_SAMPLE_FMT_U8
	// AVSampleFmtS16 wraps AV_SAMPLE_FMT_S16.
	AVSampleFmtS16 AVSampleFormat = C.AV_SAMPLE_FMT_S16
	// AVSampleFmtS32 wraps AV_SAMPLE_FMT_S32.
	AVSampleFmtS32 AVSampleFormat = C.AV_SAMPLE_FMT_S32
	// AVSampleFmtFlt wraps AV_SAMPLE_FMT_FLT.
	AVSampleFmtFlt AVSampleFormat = C.AV_SAMPLE_FMT_FLT
	// AVSampleFmtDbl wraps AV_SAMPLE_FMT_DBL.
	AVSampleFmtDbl AVSampleFormat = C.AV_SAMPLE_FMT_DBL
	// AVSampleFmtU8P wraps AV_SAMPLE_FMT_U8P.
	AVSampleFmtU8P AVSampleFormat = C.AV_SAMPLE_FMT_U8P
	// AVSampleFmtS16P wraps AV_SAMPLE_FMT_S16P.
	AVSampleFmtS16P AVSampleFormat = C.AV_SAMPLE_FMT_S16P
	// AVSampleFmtS32P wraps AV_SAMPLE_FMT_S32P.
	AVSampleFmtS32P AVSampleFormat = C.AV_SAMPLE_FMT_S32P
	// AVSampleFmtFltp wraps AV_SAMPLE_FMT_FLTP.
	AVSampleFmtFltp AVSampleFormat = C.AV_SAMPLE_FMT_FLTP
	// AVSampleFmtDblp wraps AV_SAMPLE_FMT_DBLP.
	AVSampleFmtDblp AVSampleFormat = C.AV_SAMPLE_FMT_DBLP
	// AVSampleFmtS64 wraps AV_SAMPLE_FMT_S64.
	AVSampleFmtS64 AVSampleFormat = C.AV_SAMPLE_FMT_S64
	// AVSampleFmtS64P wraps AV_SAMPLE_FMT_S64P.
	AVSampleFmtS64P AVSampleFormat = C.AV_SAMPLE_FMT_S64P
	// AVSampleFmtNb wraps AV_SAMPLE_FMT_NB.
	AVSampleFmtNb AVSampleFormat = C.AV_SAMPLE_FMT_NB
)
