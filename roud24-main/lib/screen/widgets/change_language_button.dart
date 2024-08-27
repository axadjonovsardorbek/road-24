import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:zoom_tap_animation/zoom_tap_animation.dart';
import '../../utils/color/app_colors.dart';
import '../../utils/style/app_text_style.dart';

Widget changeLanguageButton(BuildContext context){
  return   ZoomTapAnimation(
    child: DropdownButton<Locale>(
      value: context.locale,
      style: AppTextStyle.semiBold.copyWith(fontSize: 20.w),
      icon: Icon(
        Icons.arrow_forward_ios,
        color: AppColors.c356899,
        size: 20.w,
      ),
      onChanged: (Locale? locale) {
        if (locale != null) {
          debugPrint('locale code : $locale');

          context.setLocale(locale);
        }
      },
      items: context.supportedLocales.map((locale) {
        return DropdownMenuItem(
          value: locale,
          child: Text(
            locale.languageCode.toUpperCase(),
            style: AppTextStyle.semiBold
                .copyWith(fontSize: 20.w),
          ),
        );
      }).toList(),
    ),
  );
}