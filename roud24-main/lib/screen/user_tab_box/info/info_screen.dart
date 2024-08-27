import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/screen/widgets/change_language_button.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/style/app_text_style.dart';

import 'grid_item.dart';

class InfoScreen extends StatelessWidget {
  const InfoScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text(
          'info_screen_title'.tr(),
          style: AppTextStyle.bold.copyWith(
            fontSize: 24.w,
          ),
        ),
        backgroundColor: Colors.transparent,
        elevation: 0,
        actions: [changeLanguageButton(context)],
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: GridView.count(
          crossAxisCount: 2,
          mainAxisSpacing: 16,
          crossAxisSpacing: 16,
          children: [
            buildGridItem(Icons.block, 'penalties'.tr()),
            buildGridItem(Icons.car_repair, 'where_to_check'.tr()),
            buildGridItem(Icons.book, 'signs_and_rules'.tr()),
            buildGridItem(Icons.car_rental, 'buy_sell_car'.tr()),
            buildGridItem(Icons.build, 'car_modification'.tr()),
            buildGridItem(Icons.help_outline, 'faq'.tr()),
          ],
        ),
      ),
    );
  }
}
