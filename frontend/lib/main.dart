import 'package:flutter/material.dart';
import 'package:telepathy_admin/environment.dart';
import 'package:telepathy_admin/root_app.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();

  Environment.init(    
    apiBaseUrl: 'apiBaseUrl',
  );

  runApp(const RootApp());
}


