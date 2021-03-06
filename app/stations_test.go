package app

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

// http://radiko.jp/v2/station/list/JP13.xml

const MockStationsXML = `<?xml version="1.0" encoding="UTF-8" ?>
<stations area_id="JP13" area_name="TOKYO JAPAN">
 <station>
    <id>TBS</id>
    <name>TBSラジオ</name>
    <ascii_name>TBS RADIO</ascii_name>
    <href>http://www.tbs.co.jp/radio/</href>
    <logo_xsmall>http://radiko.jp/station/logo/TBS/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/TBS/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/TBS/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/TBS/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/TBS/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/TBS/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/TBS/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/TBS/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/TBS/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/TBS/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/TBS/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/TBS/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/TBS/258x60.png</logo>
<feed>http://radiko.jp/station/feed/TBS.xml</feed>
<banner>http://radiko.jp/res/banner/TBS/20130329155819.jpg</banner>
</station>
 <station>
    <id>QRR</id>
    <name>文化放送</name>
    <ascii_name>JOQR  BUNKA HOSO</ascii_name>
    <href>http://www.joqr.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/QRR/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/QRR/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/QRR/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/QRR/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/QRR/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/QRR/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/QRR/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/QRR/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/QRR/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/QRR/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/QRR/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/QRR/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/QRR/258x60.png</logo>
<feed>http://radiko.jp/station/feed/QRR.xml</feed>
<banner>http://radiko.jp/res/banner/QRR/20160719195922.png</banner>
</station>
 <station>
    <id>LFR</id>
    <name>ニッポン放送</name>
    <ascii_name>JOLF  NIPPON HOSO</ascii_name>
    <href>http://www.1242.com/</href>
    <logo_xsmall>http://radiko.jp/station/logo/LFR/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/LFR/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/LFR/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/LFR/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/LFR/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/LFR/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/LFR/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/LFR/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/LFR/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/LFR/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/LFR/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/LFR/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/LFR/258x60.png</logo>
<feed>http://radiko.jp/station/feed/LFR.xml</feed>
<banner>http://radiko.jp/res/banner/LFR/20160328113607.jpg</banner>
</station>
 <station>
    <id>RN1</id>
    <name>ラジオNIKKEI第1 </name>
    <ascii_name>RADIONIKKEI </ascii_name>
    <href>http://www.radionikkei.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/RN1/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/RN1/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/RN1/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/RN1/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/RN1/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/RN1/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/RN1/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/RN1/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/RN1/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/RN1/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/RN1/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/RN1/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/RN1/258x60.png</logo>
<feed>http://radiko.jp/station/feed/RN1.xml</feed>
<banner>http://radiko.jp/res/banner/RN1/20150708182111.png</banner>
</station>
 <station>
    <id>RN2</id>
    <name>ラジオNIKKEI第2</name>
    <ascii_name>RADIONIKKEI2</ascii_name>
    <href>http://www.radionikkei.jp/ </href>
    <logo_xsmall>http://radiko.jp/station/logo/RN2/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/RN2/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/RN2/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/RN2/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/RN2/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/RN2/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/RN2/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/RN2/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/RN2/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/RN2/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/RN2/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/RN2/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/RN2/258x60.png</logo>
<feed>http://radiko.jp/station/feed/RN2.xml</feed>
<banner>http://radiko.jp/res/banner/RN2/20151125133122.png</banner>
</station>
 <station>
    <id>INT</id>
    <name>InterFM897</name>
    <ascii_name>InterFM897</ascii_name>
    <href>http://www.interfm.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/INT/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/INT/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/INT/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/INT/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/INT/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/INT/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/INT/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/INT/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/INT/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/INT/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/INT/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/INT/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/INT/258x60.png</logo>
<feed>http://radiko.jp/station/feed/INT.xml</feed>
<banner>http://radiko.jp/res/banner/INT/20151001094947.png</banner>
</station>
 <station>
    <id>FMT</id>
    <name>TOKYO FM</name>
    <ascii_name>TOKYO FM</ascii_name>
    <href>http://www.tfm.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/FMT/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/FMT/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/FMT/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/FMT/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/FMT/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/FMT/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/FMT/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/FMT/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/FMT/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/FMT/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/FMT/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/FMT/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/FMT/258x60.png</logo>
<feed>http://radiko.jp/station/feed/FMT.xml</feed>
<banner>http://radiko.jp/res/banner/FMT/20160901113536.png</banner>
</station>
 <station>
    <id>FMJ</id>
    <name>J-WAVE</name>
    <ascii_name>J-WAVE</ascii_name>
    <href>http://www.j-wave.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/FMJ/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/FMJ/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/FMJ/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/FMJ/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/FMJ/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/FMJ/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/FMJ/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/FMJ/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/FMJ/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/FMJ/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/FMJ/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/FMJ/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/FMJ/258x60.png</logo>
<feed>http://radiko.jp/station/feed/FMJ.xml</feed>
<banner>http://radiko.jp/res/banner/FMJ/20150403151103.jpg</banner>
</station>
 <station>
    <id>JORF</id>
    <name>ラジオ日本</name>
    <ascii_name>RF RADIO NIPPON</ascii_name>
    <href>http://www.jorf.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/JORF/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/JORF/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/JORF/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/JORF/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/JORF/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/JORF/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/JORF/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/JORF/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/JORF/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/JORF/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/JORF/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/JORF/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/JORF/258x60.png</logo>
<feed>http://radiko.jp/station/feed/JORF.xml</feed>
<banner>http://radiko.jp/res/banner/JORF/20150522120802.jpg</banner>
</station>
 <station>
    <id>BAYFM78</id>
    <name>bayfm78</name>
    <ascii_name>bayfm78</ascii_name>
    <href>http://www.bayfm.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/BAYFM78/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/BAYFM78/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/BAYFM78/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/BAYFM78/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/BAYFM78/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/BAYFM78/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/BAYFM78/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/BAYFM78/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/BAYFM78/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/BAYFM78/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/BAYFM78/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/BAYFM78/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/BAYFM78/258x60.png</logo>
<feed>http://radiko.jp/station/feed/BAYFM78.xml</feed>
<banner>http://radiko.jp/res/banner/BAYFM78/20150824112330.png</banner>
</station>
 <station>
    <id>NACK5</id>
    <name>NACK5</name>
    <ascii_name>NACK5</ascii_name>
    <href>http://www.nack5.co.jp</href>
    <logo_xsmall>http://radiko.jp/station/logo/NACK5/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/NACK5/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/NACK5/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/NACK5/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/NACK5/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/NACK5/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/NACK5/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/NACK5/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/NACK5/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/NACK5/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/NACK5/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/NACK5/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/NACK5/258x60.png</logo>
<feed>http://radiko.jp/station/feed/NACK5.xml</feed>
<banner>http://radiko.jp/res/banner/NACK5/20110926101714.png</banner>
</station>
 <station>
    <id>YFM</id>
    <name>ＦＭヨコハマ</name>
    <ascii_name>Fm yokohama 84.7</ascii_name>
    <href>http://www.fmyokohama.co.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/YFM/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/YFM/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/YFM/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/YFM/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/YFM/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/YFM/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/YFM/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/YFM/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/YFM/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/YFM/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/YFM/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/YFM/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/YFM/258x60.png</logo>
<feed>http://radiko.jp/station/feed/YFM.xml</feed>
<banner>http://radiko.jp/res/banner/YFM/20110922163525.png</banner>
</station>
 <station>
    <id>HOUSOU-DAIGAKU</id>
    <name>放送大学</name>
    <ascii_name>HOUSOU-DAIGAKU</ascii_name>
    <href>http://www.ouj.ac.jp/</href>
    <logo_xsmall>http://radiko.jp/station/logo/HOUSOU-DAIGAKU/logo_xsmall.png</logo_xsmall>
    <logo_small>http://radiko.jp/station/logo/HOUSOU-DAIGAKU/logo_small.png</logo_small>
    <logo_medium>http://radiko.jp/station/logo/HOUSOU-DAIGAKU/logo_medium.png</logo_medium>
    <logo_large>http://radiko.jp/station/logo/HOUSOU-DAIGAKU/logo_large.png</logo_large>
<logo width="124" height="40">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/124x40.png</logo>
<logo width="344" height="80">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/344x80.png</logo>
<logo width="688" height="160">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/688x160.png</logo>
<logo width="172" height="40">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/172x40.png</logo>
<logo width="224" height="100">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/224x100.png</logo>
<logo width="448" height="200">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/448x200.png</logo>
<logo width="112" height="50">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/112x50.png</logo>
<logo width="168" height="75">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/168x75.png</logo>
<logo width="258" height="60">http://radiko.jp/v2/static/station/logo/HOUSOU-DAIGAKU/258x60.png</logo>
<feed>http://radiko.jp/station/feed/HOUSOU-DAIGAKU.xml</feed>
<banner>http://radiko.jp/res/banner/HOUSOU-DAIGAKU/20150805145127.png</banner>
</station>
</stations>
`

func _TestStations(t *testing.T, stations []Station) {
	if len(stations) != 13 {
		t.Errorf("Expected %d but got %d", 13, len(stations))
	}
	if stations[0].ID != "TBS" {
		t.Errorf("Expected %s but got %s", "TBS", stations[0].ID)
	}
	if stations[0].Name != "TBSラジオ" {
		t.Errorf("Expected %s but got %s", "TBSラジオ", stations[0].Name)
	}
	if stations[0].AsciiName != "TBS RADIO" {
		t.Errorf("Expected %s but got %s", "TBS RADIO", stations[0].AsciiName)
	}
	if stations[0].Href != "http://www.tbs.co.jp/radio/" {
		t.Errorf("Expected %s but got %s", "http://www.tbs.co.jp/radio/", stations[0].Href)
	}
}

func _SetupMockHTTP() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://radiko.jp/v2/station/list/JP13.xml",
		httpmock.NewStringResponder(200, MockStationsXML))
}

func TestGetStations(t *testing.T) {
	stations := GetStations([]byte(MockStationsXML))
	_TestStations(t, stations)
}

func TestFetchStations(t *testing.T) {
	_SetupMockHTTP()
	stations, err := FetchStations("東京")
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	_TestStations(t, stations)
}
