//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class UpsertOryAccessControlPolicyRoleOK {
  /// Returns a new [UpsertOryAccessControlPolicyRoleOK] instance.
  UpsertOryAccessControlPolicyRoleOK({
    this.payload,
  });

  OryAccessControlPolicyRole payload;

  @override
  bool operator ==(Object other) => identical(this, other) || other is UpsertOryAccessControlPolicyRoleOK &&
     other.payload == payload;

  @override
  int get hashCode =>
    (payload == null ? 0 : payload.hashCode);

  @override
  String toString() => 'UpsertOryAccessControlPolicyRoleOK[payload=$payload]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (payload != null) {
      json[r'Payload'] = payload;
    }
    return json;
  }

  /// Returns a new [UpsertOryAccessControlPolicyRoleOK] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static UpsertOryAccessControlPolicyRoleOK fromJson(Map<String, dynamic> json) => json == null
    ? null
    : UpsertOryAccessControlPolicyRoleOK(
        payload: OryAccessControlPolicyRole.fromJson(json[r'Payload']),
    );

  static List<UpsertOryAccessControlPolicyRoleOK> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <UpsertOryAccessControlPolicyRoleOK>[]
      : json.map((v) => UpsertOryAccessControlPolicyRoleOK.fromJson(v)).toList(growable: true == growable);

  static Map<String, UpsertOryAccessControlPolicyRoleOK> mapFromJson(Map<String, dynamic> json) {
    final map = <String, UpsertOryAccessControlPolicyRoleOK>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = UpsertOryAccessControlPolicyRoleOK.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of UpsertOryAccessControlPolicyRoleOK-objects as value to a dart map
  static Map<String, List<UpsertOryAccessControlPolicyRoleOK>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<UpsertOryAccessControlPolicyRoleOK>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = UpsertOryAccessControlPolicyRoleOK.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

