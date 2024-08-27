import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:roud24/utils/extension/extension.dart';

import '../../../utils/color/app_colors.dart';
import '../../../utils/style/app_text_style.dart';

class InputUserText extends StatefulWidget {
  const InputUserText({
    super.key,
    required this.controller,
    this.isIcon = false,
    required this.title,
    this.errorText = 'example@gmail.com',
    this.isRightBorder = true,
    this.maxLength = 2,
    this.borderAll = false,
    this.questionOnTap,
    this.isQuestion = false,
  });

  final TextEditingController controller;
  final bool isIcon, isRightBorder, borderAll, isQuestion;
  final String errorText, title;
  final int maxLength;
  final VoidCallback? questionOnTap;

  @override
  State<InputUserText> createState() => _InputUserTextState();
}

class _InputUserTextState extends State<InputUserText> {
  @override
  Widget build(BuildContext context) {
    return TextFormField(
      maxLength: widget.maxLength,
      controller: widget.controller,
      scrollPadding: EdgeInsets.symmetric(horizontal: 10.w),
      style: AppTextStyle.bold.copyWith(
        fontSize: 20.w,
      ),
      textCapitalization: TextCapitalization.characters,
      decoration: InputDecoration(
        hintText: widget.title.tr(),
        hintStyle: AppTextStyle.bold.copyWith(
          fontSize: 20.w,
        ),
        suffixIcon: widget.isIcon
            ? RichText(
                text: TextSpan(
                    text: "🇺🇿",
                    style: AppTextStyle.bold.copyWith(fontSize: 25.w),
                    children: [
                      TextSpan(
                        text: '\nUZ',
                        style: AppTextStyle.bold.copyWith(
                          color: Colors.blue,
                          fontSize: 12.w,
                        ),
                      )
                    ]),
              )
            : widget.isQuestion
                ? IconButton(
                  onPressed: widget.questionOnTap,
                  icon: Container(
                    decoration: BoxDecoration(
                        border: Border.all(
                          color: Colors.red,
                          width: 2,
                        ),
                        shape: BoxShape.circle
                    ),
                    child: const Icon(
                      Icons.question_mark_outlined,
                      size: 24,
                      color: Colors.red,
                    ),
                  ),
                )
                : null,
        disabledBorder: OutlineInputBorder(
          borderRadius: getBorderRadius(
            isRightBorder: widget.isRightBorder,
            borderAll: widget.borderAll,
          ),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: getBorderRadius(
            isRightBorder: widget.isRightBorder,
            borderAll: widget.borderAll,
          ),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: getBorderRadius(
            isRightBorder: widget.isRightBorder,
            borderAll: widget.borderAll,
          ),
          borderSide: const BorderSide(
            width: 1,
            color: AppColors.c95969D,
          ),
        ),
        errorBorder: OutlineInputBorder(
          borderRadius: getBorderRadius(
            isRightBorder: widget.isRightBorder,
            borderAll: widget.borderAll,
          ),
          borderSide: const BorderSide(
            width: 1,
            color: Colors.red,
          ),
        ),
        focusedErrorBorder: OutlineInputBorder(
          borderRadius: getBorderRadius(
            isRightBorder: widget.isRightBorder,
            borderAll: widget.borderAll,
          ),
          borderSide: const BorderSide(
            width: 1,
            color: Colors.red,
          ),
        ),
      ),
    );
  }
}

BorderRadius getBorderRadius({
  required bool isRightBorder,
  required bool borderAll,
}) {
  return borderAll
      ? BorderRadius.circular(16.w)
      : isRightBorder
          ? BorderRadius.only(
              topLeft: Radius.circular(16.w),
              bottomLeft: Radius.circular(16.w),
            )
          : BorderRadius.only(
              topRight: Radius.circular(16.w),
              bottomRight: Radius.circular(16.w),
            );
}
