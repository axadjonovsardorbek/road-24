import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/style/app_text_style.dart';
import 'package:zoom_tap_animation/zoom_tap_animation.dart';

Widget globalButton({
  required String title,
  required VoidCallback onTap,
  Color? background, textColor, loadingColor, borderColor,
  TextStyle? textStyle,
  bool isLoading = false,
  double? borderRadius, buttonW, buttonH,
}) {
  return ZoomTapAnimation(
    onTap: onTap,
    child: Container(
      width: buttonW ?? 70.w,
      height: buttonH ?? 50.h,
      decoration: BoxDecoration(
        color: background ?? Colors.blueAccent,
        borderRadius: BorderRadius.circular(borderRadius ?? 16.w),
        border: Border.all(
          width: 1,
          color: borderColor ?? Colors.blueAccent,
        ),
      ),
      child: Center(
        child: isLoading? CupertinoActivityIndicator(color: loadingColor ?? Colors.white,):Text(
          title,
          style: textStyle ?? AppTextStyle.bold.copyWith(
            color: textColor ?? Colors.white,
          ),
          textAlign: TextAlign.center,
        ),
      ),
    ),
  );
}
