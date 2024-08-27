import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/utils/extension/extension.dart';
import '../../../utils/color/app_colors.dart';
import '../../../utils/style/app_text_style.dart';

class InputText extends StatefulWidget {
  const InputText({
    super.key,
    required this.controller,
    this.isPassword = false,
    required this.title,
    required this.regExp,
    this.errorText = 'example@gmail.com',
  });

  final TextEditingController controller;
  final bool isPassword;
  final String title;
  final RegExp regExp;
  final String errorText;

  @override
  State<InputText> createState() => _InputTextState();
}

class _InputTextState extends State<InputText> {
  bool isVisible = true;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: widget.controller,
      autovalidateMode: AutovalidateMode.onUserInteraction,
      scrollPadding: EdgeInsets.symmetric(horizontal: 10.w),
      style: AppTextStyle.regular,
      obscureText: widget.isPassword ? isVisible : false,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return widget.errorText;
        }
          if (!widget.regExp!.hasMatch(value)) {
            return widget.errorText;
        }
        return null;
      },
      decoration: InputDecoration(
        hintText: widget.title.tr(),
        hintStyle: AppTextStyle.regular,
        suffixIcon: widget.isPassword
            ? IconButton(
          onPressed: () {
            setState(() {
              isVisible = !isVisible;
            },);
          },
          icon: Icon(
            isVisible ? Icons.visibility_off : Icons.remove_red_eye,
          ),
        )
            : null,
        disabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12.w),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12.w),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12.w),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),),
        errorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12.w),
          borderSide: const BorderSide(
            width: 1,
            color: Colors.red,
          ),
        ),
        focusedErrorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12.w),
          borderSide: const BorderSide(
            width: 1,
            color: Colors.red,
          ),
        ),
      ),
    );
  }
}