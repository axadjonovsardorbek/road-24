
import 'package:flutter/material.dart';

double width = 0.0;
double height = 0.0;

extension Size on num {
  //SizedBox height
  SizedBox boxH() => SizedBox(height: (this / 812) * height);
  //SizedBox width
  SizedBox boxW() => SizedBox(width: (this / 375) * width);

  double get h => (this / 812) * height;

  double get w => (this / 375) * width;
}

//MediaQuery size
extension ContextExtension on BuildContext {
  double getWidth() {
    return MediaQuery.of(this).size.width;
  }

  double getHeight() {
    double appBarHeight = AppBar().preferredSize.height;
    double statusBarHeight = MediaQuery.of(this).viewPadding.top;
    double screenHeight =
        MediaQuery.of(this).size.height - appBarHeight - statusBarHeight;
    return screenHeight;
  }
}


extension PaddingExtension on Widget {
  // Barcha tomonlar uchun bir xil padding
  Widget paddingAll(double value) => Padding(
    padding: EdgeInsets.all(value),
    child: this,
  );

  // Faqat yuqori qism uchun padding
  Widget paddingTop(double value) => Padding(
    padding: EdgeInsets.only(top: value),
    child: this,
  );

  // Faqat pastki qism uchun padding
  Widget paddingBottom(double value) => Padding(
    padding: EdgeInsets.only(bottom: value),
    child: this,
  );

  // Faqat chap qism uchun padding
  Widget paddingLeft(double value) => Padding(
    padding: EdgeInsets.only(left: value),
    child: this,
  );

  // Faqat o'ng qism uchun padding
  Widget paddingRight(double value) => Padding(
    padding: EdgeInsets.only(right: value),
    child: this,
  );

  // Gorizontal (chap va o'ng) uchun padding
  Widget paddingHorizontal(double value) => Padding(
    padding: EdgeInsets.symmetric(horizontal: value),
    child: this,
  );

  // Vertikal (yuqori va pastki) uchun padding
  Widget paddingVertical(double value) => Padding(
    padding: EdgeInsets.symmetric(vertical: value),
    child: this,
  );

  // To'liq moslashuvchan padding
  Widget padding({
    double top = 0,
    double bottom = 0,
    double left = 0,
    double right = 0,
  }) =>
      Padding(
        padding: EdgeInsets.only(
          top: top,
          bottom: bottom,
          left: left,
          right: right,
        ),
        child: this,
      );
}