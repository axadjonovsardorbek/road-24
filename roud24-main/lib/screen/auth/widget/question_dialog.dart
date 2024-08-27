import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/screen/widgets/global_button.dart';
import 'package:roud24/screen/widgets/important_know_screen.dart';
import 'package:roud24/screen/widgets/know_button.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/images/app_images.dart';
import 'package:roud24/utils/style/app_text_style.dart';

Widget questionDialog(BuildContext context) {
  return AlertDialog(
    title: SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(
            'certificate'.tr(),
            style: AppTextStyle.bold,
          ),
          Image.asset(
            AppImages.idOne,
            width: 270.w,
            height: 200.h,
            fit: BoxFit.contain,
          ),
          Image.asset(
            AppImages.idTwo,
            width: 270.w,
            height: 100.h,
          ),
          10.boxH(),
          knowButton(
            buttonW: 15.0,
            buttonH: 10.0,
            title: "know".tr(),
            onTap: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (_) => const ImportantKnowScreen(),
                ),
              );
            },
          ),
          10.boxH(),
          globalButton(
            buttonW: double.infinity,
            title: "back".tr(),
            onTap: () {
              Navigator.pop(context);
            },
          )
        ],
      ),
    ),
  );
}
