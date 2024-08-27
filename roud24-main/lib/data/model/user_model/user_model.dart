import 'package:roud24/data/model/image_model/image_model.dart';

import 'jarima/jarima_model.dart';

class UserModel {
  final String carNumber;
  final String color;
  final ImageModel imageModel;
  final List<JarimaModel> jarimalar;
  final String textPasswordDate;
  final String straxovka;

  UserModel({
    required this.carNumber,
    required this.color,
    required this.imageModel,
    required this.jarimalar,
    required this.textPasswordDate,
    required this.straxovka,
  });

  // From JSON
  factory UserModel.fromJson(Map<String, dynamic> json) => UserModel(
    carNumber: json['car_number'] as String? ?? '',
    color: json['color'] as String? ?? '',
    imageModel: ImageModel.fromJson(json['image_model'] as Map<String, dynamic>? ?? {}),
    jarimalar: (json['jarimalar'] as List<dynamic>?)
        ?.map((e) => JarimaModel.fromJson(e as Map<String, dynamic>))
        .toList() ??
        [],
    textPasswordDate: json['text_password_date'] as String? ?? '',
    straxovka: json['straxovka'] as String? ?? '',
  );

  // To JSON
  Map<String, dynamic> toJson() => {
    "car_number": carNumber,
    "color": color,
    "image_model": imageModel.toJson(),
    "jarimalar": jarimalar.map((e) => e.toJson()).toList(),
    "text_password_date": textPasswordDate,
    "straxovka": straxovka,
  };

  // Copy with method
  UserModel copyWith({
    String? carNumber,
    String? color,
    ImageModel? imageModel,
    List<JarimaModel>? jarimalar,
    String? textPasswordDate,
    String? straxovka,
  }) =>
      UserModel(
        carNumber: carNumber ?? this.carNumber,
        color: color ?? this.color,
        imageModel: imageModel ?? this.imageModel,
        jarimalar: jarimalar ?? this.jarimalar,
        textPasswordDate: textPasswordDate ?? this.textPasswordDate,
        straxovka: straxovka ?? this.straxovka,
      );

  // Initial value
  static UserModel initialValue() => UserModel(
    carNumber: '',
    color: '',
    imageModel: ImageModel.initialValue(),
    jarimalar: [],
    textPasswordDate: '',
    straxovka: '',
  );
}
