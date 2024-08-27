import 'package:cached_network_image/cached_network_image.dart';
import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';
import 'package:roud24/cubit/image_cubit/image_cubit.dart';
import 'package:roud24/cubit/image_cubit/image_state.dart';
import 'package:roud24/data/model/forms_status.dart';
import 'package:roud24/data/model/image_model/image_model.dart';
import 'package:roud24/screen/auth/widget/input_user_text.dart';
import 'package:roud24/screen/auth/widget/question_dialog.dart';
import 'package:roud24/screen/route.dart';
import 'package:roud24/screen/widgets/change_language_button.dart';
import 'package:roud24/screen/widgets/global_button.dart';
import 'package:roud24/screen/widgets/global_laoding_button.dart';
import 'package:roud24/utils/color/app_colors.dart';
import 'package:roud24/utils/extension/extension.dart';
import 'package:roud24/utils/style/app_text_style.dart';
import 'package:roud24/utils/ui_utils/ui_utils.dart';
import 'package:shimmer/shimmer.dart';

class UserLoginScreen extends StatefulWidget {
  const UserLoginScreen({super.key});

  @override
  State<UserLoginScreen> createState() => _UserLoginScreenState();
}

class _UserLoginScreenState extends State<UserLoginScreen> {
  final TextEditingController numberOneController = TextEditingController();
  final TextEditingController numberTwoController = TextEditingController();
  String? _selectedCarType;
  ImageModel globalImageModel = ImageModel.initialValue();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            changeLanguageButton(context),
            20.boxH(),
            Text(
              "auto_number".tr(),
              style: AppTextStyle.bold,
            ),
            5.boxH(),
            Row(
              children: [
                Expanded(
                  flex: 2,
                  child: InputUserText(
                    controller: numberOneController,
                    title: "01",
                  ),
                ),
                Expanded(
                  flex: 8,
                  child: InputUserText(
                    controller: numberTwoController,
                    title: "A 111 AAA",
                    isRightBorder: false,
                    isIcon: true,
                    maxLength: 8,
                  ),
                ),
              ],
            ),
            10.boxH(),
            Text(
              "passport_number".tr(),
              style: AppTextStyle.bold,
            ),
            5.boxH(),
            Row(
              children: [
                Expanded(
                  flex: 3,
                  child: InputUserText(
                    controller: numberOneController,
                    title: "AA",
                    borderAll: true,
                    maxLength: 2,
                  ),
                ),
                20.boxW(),
                Expanded(
                  flex: 7,
                  child: InputUserText(
                    controller: numberTwoController,
                    title: "123 45 67",
                    isQuestion: true,
                    borderAll: true,
                    maxLength: 7,
                    questionOnTap: () {
                      showDialog(
                        context: context,
                        builder: (context) {
                          return questionDialog(context);
                        },
                      );
                    },
                  ),
                ),
              ],
            ),
            10.boxH(),
            Text(
              'car_type'.tr(),
              style: AppTextStyle.bold,
            ),
            10.boxH(),
            BlocBuilder<ImageCubit, ImageState>(
              builder: (BuildContext context, ImageState state) {
                if (state.status == FormsStatus.fetch ) {
                  return Shimmer.fromColors(
                    baseColor: Colors.grey[300]!,
                    highlightColor: Colors.grey[100]!,
                    child: Container(
                      width: double.infinity,
                      height: 48.0,
                      decoration: BoxDecoration(
                        color: Colors.grey[300],
                        borderRadius: BorderRadius.circular(5),
                      ),
                    ),
                  );
                }
                return DropdownButton<String>(
                  value: _selectedCarType,
                  hint: const Text('Select Car Type'),
                  items: state.listImages.map((imageModel) {
                    return DropdownMenuItem<String>(
                      value: imageModel.carType,
                      child: Text(imageModel.carType),
                    );
                  }).toList(),
                  onChanged: (String? newValue) {
                    setState(() {
                      _selectedCarType = newValue;
                    });
                  },
                );
              },
            ),
            BlocBuilder<ImageCubit, ImageState>(
              builder: (BuildContext context, ImageState state) {
                return Column(
                  children: [
                    if (state.imagesModel.imagePath.isNotEmpty)
                      GestureDetector(
                        onTap: () {
                          showErrorMessage(
                            message: "delete",
                            context: context,
                            onTap: () {
                              context.read<ImageCubit>().deleteImage();
                              Navigator.pop(context);
                            },
                            isDelete: true,
                          );
                        },
                        child: CachedNetworkImage(
                          width: double.infinity,
                          height: 200.h,
                          fit: BoxFit.cover,
                          imageUrl: state.imagesModel.imagePath,
                          placeholder: (context, url) =>
                              const CupertinoActivityIndicator(
                            color: AppColors.c000000,
                          ),
                        ),
                      ),
                    10.boxH(),
                    LoginButton(
                      title: "image".tr(),
                      isLoading: state.status == FormsStatus.loading,
                      onTap: () {
                        if (_selectedCarType != null) {
                          context.read<ImageCubit>().insertImage(
                                source: ImageSource.gallery,
                                imageModel: ImageModel(
                                  carType: _selectedCarType!,
                                  imagePath: '',
                                ),
                              );
                        } else {
                          showErrorMessage(
                            message: "null_car_type",
                            context: context,
                            onTap: () {
                              Navigator.pop(context);
                            },
                          );
                        }
                      },
                    ),
                  ],
                );
              },
            ),
            20.boxH(),
            globalButton(
              buttonW: double.infinity,
              title: 'login'.tr(),
              onTap: () {
                Navigator.pushNamed(context, RouteName.tabBox);
              },
            )
          ],
        ).paddingAll(30.w),
      ),
    );
  }

  @override
  void dispose() {
    numberOneController.dispose();
    numberTwoController.dispose();
    super.dispose();
  }
}
