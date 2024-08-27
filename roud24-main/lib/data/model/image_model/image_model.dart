class ImageModel {
  final String carType;
  final String imagePath;

  ImageModel({
    required this.carType,
    required this.imagePath,
  });

  factory ImageModel.fromJson(Map<String, dynamic> json) => ImageModel(
        carType: json['car_type'] as String? ?? '',
        imagePath: json['image_path'] as String? ?? '',
      );

  Map<String, dynamic> toJson() => {
        "car_type": carType,
        "image_path": imagePath,
      };

  ImageModel copyWith({
    String? imagePath,
    String? carType,
  }) =>
      ImageModel(
        carType: carType ?? this.carType,
        imagePath: imagePath ?? this.imagePath,
      );

  static ImageModel initialValue()=>ImageModel(carType: '', imagePath: '');
}
