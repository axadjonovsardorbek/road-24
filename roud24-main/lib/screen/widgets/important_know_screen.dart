import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:flutter_widget_from_html/flutter_widget_from_html.dart';
import 'package:roud24/utils/color/app_colors.dart';

class ImportantKnowScreen extends StatelessWidget {
  const ImportantKnowScreen({super.key});

  @override
  Widget build(BuildContext context) {
    const String htmlContent = '''
    <p><strong>Bilish Muhum</strong></p>
    <p>Esk namunadagi haydovchilar guvohnomalari va texpasportlari yangi turdagilariga majburiy almashtirish muddatlari.</p>
    <ul>
      <li><strong>2010 - yilgacha berilganlar uchun</strong> - <span style="color: red;">(2023 yil 31 mart)</span> gacha</li>
      <li><strong>2010-2012 - yilgacha berilganlar uchun</strong> - <span style="color: red;">(2023 yil 30 iyun)</span> gacha</li>
      <li><strong>2013-2015 - yilgacha berilganlar uchun</strong> - <span style="color: red;">(2023 yil 30 sentyabr)</span> gacha</li>
      <li><strong>2015 - yilgacha berilganlar uchun</strong> - <span style="color: red;">(2023 yil 31 dekabr)</span> gacha</li>
    </ul>
    <p><strong>ALMASHTIRISH KECHIKKANI UCHUN/GUVOHNOMASIZ BOSHQARISH UCHUN JARIMA</strong></p>
    <p><strong>135-modda.</strong> Yo'l harakati qoidalarida nazarda tutilgan hujjatlar bo'lmagan shaxslar tomonidan transport vositalari boshqarish; Jarima miqdori-<span style="color: red;">1 BHM (330 000 so'm)</span>.</p>
    <p><strong>TEPASPORT VA HAYDOVCHILIK GUVOHNOMASINI ALMASHTIRISH NARXI</strong></p>
    <ul>
      <li>Texpasportni yangi namunadagisini yangi namunadagisiga almashtirish xizmatining narxi <span style="color: red;">0. BHM (231 000 so'm)</span></li>
      <li>Haydovchini guvohnomasini yangi namunadagisini yangi namunadagisiga almashtirish xizmatining narxi <span style="color: red;">0. BHM (231 000 so'm)</span></li>
    </ul>
    ''';

    return Scaffold(
      appBar: AppBar(
        scrolledUnderElevation: 0,
        automaticallyImplyLeading: false,
        leading: IconButton(
          onPressed: () {
            Navigator.pop(context);
          },
          icon: const Icon(
            Icons.arrow_back_ios,
            size: 24,
            color: AppColors.c000000,
          ),
        ),
        title: Text(
          'know'.tr(),
        ),
      ),
      body: const Padding(
        padding: EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: HtmlWidget(
            htmlContent,
            textStyle: TextStyle(fontSize: 16.0),
          ),
        ),
      ),
    );
  }
}
