import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:roud24/screen/auth/user/user_login_screen.dart';
import 'package:roud24/screen/auth/ypx/ypx_login_screen.dart';
import 'package:roud24/screen/splash/splash_screen.dart';
import 'package:roud24/screen/user_tab_box/user_tab_box.dart';
import 'package:roud24/utils/style/app_text_style.dart';

class AppRoute {
  static Route generateRoute(RouteSettings settings) {
    switch (settings.name) {
      case RouteName.splash:
        return navigate(const SplashScreen());

      case RouteName.loginUser:
        return navigate(const UserLoginScreen());

      case RouteName.loginYpx:
        return navigate(const YpxLoginScreen());

      case RouteName.tabBox:
        return navigate(const TabBoxScreen());

      default:
        return navigate(
          Scaffold(
            body: Center(
              child: Text(
                "Default",
                style: AppTextStyle.bold,
              ),
            ),
          ),
        );
    }
  }

  static navigate(Widget widget) =>
      CupertinoPageRoute(builder: (context) => widget);
}

class RouteName {
  static const String splash = "/";
  static const String loginYpx = "/login_ypx";
  static const String loginUser = "/login_user";
  static const String tabBox = "/tab_box";
}
