import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/main.dart';
import 'package:roud24/screen/route.dart';
import 'package:roud24/screen/widgets/change_language_button.dart';
import 'package:roud24/screen/widgets/global_button.dart';
import 'package:roud24/utils/color/app_colors.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/images/app_images.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) {
    debugPrint("_______________________________________ splash screen");
    return Scaffold(
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          100.boxH(),
          changeLanguageButton(context),
          50.boxH(),
          Center(
            child: ClipOval(
              child: Image.asset(
                AppImages.roud,
                fit: BoxFit.cover,
                width: 200.w,
                height: 200.h,
              ),
            ),
          ),
          20.boxH(),
          Center(
            child: globalButton(
              title: "user".tr(),
              onTap: () {
                Navigator.pushNamed(context, RouteName.loginUser);
              },
              buttonH: 56.h,
              buttonW: 200.w
            ),
          ),
          20.boxH(),
          Center(
            child: globalButton(
                title: "ypx".tr(),
                onTap: () {
                  Navigator.pushNamed(context, RouteName.loginYpx);
                },
                buttonH: 56.h,
                buttonW: 200.w,
                background: Colors.transparent,
                borderColor: AppColors.c000000,
                textColor: AppColors.c000000,
            ),
          ),
        ],
      ).paddingAll(14.w),
    );
  }
}
