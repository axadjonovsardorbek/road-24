class JarimaModel {
  final String cash;
  final String title;

  JarimaModel({
    required this.cash,
    required this.title,
  });

  // From JSON
  factory JarimaModel.fromJson(Map<String, dynamic> json) => JarimaModel(
    cash: json['cash'] as String? ?? '',
    title: json['title'] as String? ?? '',
  );

  // To JSON
  Map<String, dynamic> toJson() => {
    "cash": cash,
    "title": title,
  };

  // Copy with method
  JarimaModel copyWith({
    String? cash,
    String? title,
  }) =>
      JarimaModel(
        cash: cash ?? this.cash,
        title: title ?? this.title,
      );

  // Initial value
  static JarimaModel initialValue() => JarimaModel(
    cash: '',
    title: '',
  );
}
