import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:roud24/utils/extension/extension.dart';
import '../../../utils/color/app_colors.dart';
import '../../../utils/constants/constants.dart';
import '../../../utils/style/app_text_style.dart';
import '../../widgets/change_language_button.dart';
import '../../widgets/global_laoding_button.dart';
import '../widget/input_text.dart';

class YpxLoginScreen extends StatefulWidget {
  const YpxLoginScreen({super.key});

  @override
  State<YpxLoginScreen> createState() => _YpxLoginScreenState();
}

class _YpxLoginScreenState extends State<YpxLoginScreen> {

  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();



  @override
  Widget build(BuildContext context) {
    debugPrint("____________________________________----build run login");
    return AnnotatedRegion(
      value: SystemUiOverlayStyle(
        statusBarIconBrightness: Brightness.light,
        statusBarBrightness: Brightness.light,
        statusBarColor: AppColors.c356899.withOpacity(0.6),
      ),
      child: Scaffold(
        body: Padding(
          padding:  EdgeInsets.only(top: 50.h, left: 10.w, right: 10.w),
          child: SingleChildScrollView(
            child: Form(
              key: _formKey,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Align(
                    alignment: Alignment.topLeft,
                    child: changeLanguageButton(context),
                  ),
                  Center(
                    child: Text(
                      'Welcome Back',
                      style: AppTextStyle.semiBold.copyWith(
                        fontSize: 30.w,
                      ),
                      textAlign: TextAlign.center,
                    ),
                  ),
                  11.boxH(),
                  Text(
                    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor',
                    style: AppTextStyle.regular.copyWith(
                      fontSize: 12.w,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  64.boxH(),
                  Text(
                    'Login',
                    style: AppTextStyle.semiBold,
                  ),
                  10.boxH(),
                  InputText(
                    controller: _emailController,
                    title: 'Login',
                    regExp: AppConstants.emailRegExp,
                    errorText: "login_error".tr(),
                  ),
                  15.boxH(),
                  Text(
                    'password'.tr(),
                    style: AppTextStyle.semiBold,
                  ),
                  10.boxH(),
                  InputText(
                    controller: _passwordController,
                    title: 'password'.tr(),
                    regExp: AppConstants.passwordRegExp,
                    isPassword: true,
                    errorText: "password_error".tr(),
                  ),
                  30.boxH(),
                  LoginButton(
                    title: 'login'.tr(),
                    onTap:(){},
                    isLoading: false,
                  ),
                  16.boxH(),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }
}