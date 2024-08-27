import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:roud24/utils/extension/extension.dart';

import '../../utils/style/app_text_style.dart';

class LoginButton extends StatefulWidget {
  const LoginButton({
    super.key,
    required this.onTap,
    required this.isLoading,
    required this.title,
  });

  final VoidCallback onTap;
  final bool isLoading ;
  final String title;

  @override
  State<LoginButton> createState() => _LoginButtonState();
}

class _LoginButtonState extends State<LoginButton> {

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: widget.onTap,
      child: Center(
        child: AnimatedContainer(
          decoration: BoxDecoration(
            color: Colors.blueAccent,
            borderRadius: widget.isLoading?null:BorderRadius.circular(16.w),
          ),
          duration: const Duration(seconds: 1),
          curve: Curves.linear,
          width: widget.isLoading ? 70.h : 300.w,
          height: 55.h,
          child: Center(
            child: widget.isLoading
                ? const CupertinoActivityIndicator(color: Colors.white)
                : Text(
              widget.title,
              style: AppTextStyle.bold.copyWith(
                color: Colors.white,
                fontSize: 18.w,
              ),
            ),
          ),
        ),
      ),
    );
  }
}