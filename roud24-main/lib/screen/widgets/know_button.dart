import 'package:flutter/material.dart';
import 'package:roud24/utils/color/app_colors.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/style/app_text_style.dart';
import 'package:zoom_tap_animation/zoom_tap_animation.dart';

Widget knowButton({
  required String title,
  required VoidCallback onTap,
  double? borderRadius, buttonW, buttonH,
}) {
  return ZoomTapAnimation(
    onTap: onTap,
    child: Container(
      padding: EdgeInsets.symmetric(horizontal: buttonW, vertical: buttonH),
      decoration: BoxDecoration(
        color: Colors.transparent,
        borderRadius: BorderRadius.circular(borderRadius ?? 16.w),
        border: Border.all(
          width: 1,
          color: AppColors.c000000,
        ),
      ),
      child: Center(
        child: Row(
          children: [
            Container(
              decoration:const BoxDecoration(
                color: AppColors.c000000,
                shape: BoxShape.circle
              ),
              child: const Icon(
                Icons.info_outline,
                color: Colors.white,
                size: 24,
              ),
            ),
            10.boxW(),
            Text(
              title,
              style: AppTextStyle.bold,
              textAlign: TextAlign.center,
            ),
          ],
        ),
      ),
    ),
  );
}
