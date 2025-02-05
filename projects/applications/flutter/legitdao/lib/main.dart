import 'package:flutter/material.dart';
import 'package:easy_localization/easy_localization.dart';
import 'widgets/main/main_layout.dart';
import 'pages/home_page.dart';
import 'pages/about_page.dart';
import 'pages/contact_page.dart';
import 'pages/referrals_page.dart';
import 'pages/tokens_page.dart';
import 'package:reown_appkit/reown_appkit.dart';
import 'package:provider/provider.dart';
import 'pages/token_page.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await EasyLocalization.ensureInitialized();

  runApp(
    ChangeNotifierProvider(
        create: (_) => ThemeProvider(),
        child: EasyLocalization(
          supportedLocales: const [Locale('en'), Locale('fr')],
          path: 'lib/localization',
          fallbackLocale: const Locale('en'),
          child: MyApp(),
        )),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    final themeProvider = Provider.of<ThemeProvider>(context);

    final MaterialPageRoute notFoundPage = MaterialPageRoute(
      builder: (context) => Scaffold(
        body: Center(
          child: MainLayout(
            child: Center(
              // Centering the text horizontally and vertically
              child: Text(
                '404 - Wallet Address Not Found',
                style: Theme.of(context).textTheme.headlineLarge,
              ),
            ),
            isDark: themeProvider.isDark,
            onThemeToggle: () {
              themeProvider.toggleTheme();
            },
          ),
        ),
      ),
    );

    ReownAppKitModalTheme theme = ReownAppKitModalTheme(
      isDarkMode: themeProvider.isDark,
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        locale: context.locale,
        supportedLocales: context.supportedLocales,
        localizationsDelegates: context.localizationDelegates,
        theme: themeProvider.getTheme(),
        initialRoute: '/',
        onGenerateRoute: (settings) {
          final uri = Uri.parse(settings.name ?? '');

          // Handle dynamic route `/referrals/:walletAddress`
          if (uri.pathSegments.length == 2 &&
              uri.pathSegments.first == 'referrals') {
            final walletAddress = uri.pathSegments[1];

            // Check if walletAddress is null or empty
            if (walletAddress.isEmpty) {
              return notFoundPage;
            }

            return MaterialPageRoute(
              builder: (context) => MainLayout(
                child: ReferralsPage(walletAddress: walletAddress),
                isDark: themeProvider.isDark,
                onThemeToggle: () {
                  themeProvider.toggleTheme();
                },
              ),
            );
          }

          // Handle dynamic route `/tokens/:tokenSlug`
          if (uri.pathSegments.length == 2 &&
              uri.pathSegments.first == 'tokens') {
            final tokenSlug = uri.pathSegments[1];

            // Check if tokenSlug is null or empty
            if (tokenSlug.isEmpty) {
              return notFoundPage;
            }

            return MaterialPageRoute(
              builder: (context) => MainLayout(
                child: TokenPage(
                    isDark: themeProvider._isDark, tokenSlug: tokenSlug),
                isDark: themeProvider.isDark,
                onThemeToggle: () {
                  themeProvider.toggleTheme();
                },
              ),
            );
          }

          // Default routes
          switch (settings.name) {
            case '/':
              return MaterialPageRoute(
                builder: (context) => MainLayout(
                  child: HomePage(
                    isDark: themeProvider.isDark,
                  ),
                  isDark: themeProvider.isDark,
                  onThemeToggle: () {
                    themeProvider.toggleTheme();
                  },
                ),
              );
            case '/tokens':
              return MaterialPageRoute(
                builder: (context) => MainLayout(
                  child: TokensPage(
                    isDark: themeProvider.isDark,
                  ),
                  isDark: themeProvider.isDark,
                  onThemeToggle: () {
                    themeProvider.toggleTheme();
                  },
                ),
              );
            case '/about':
              return MaterialPageRoute(
                builder: (context) => MainLayout(
                  child: AboutPage(),
                  isDark: themeProvider.isDark,
                  onThemeToggle: () {
                    themeProvider.toggleTheme();
                  },
                ),
              );
            case '/contact':
              return MaterialPageRoute(
                builder: (context) => MainLayout(
                  child: ContactPage(),
                  isDark: themeProvider.isDark,
                  onThemeToggle: () {
                    themeProvider.toggleTheme();
                  },
                ),
              );
            default:
              return notFoundPage;
          }
        },
      ),
    );

    return theme;
  }
}

class ThemeProvider extends ChangeNotifier {
  bool _isDark = true;

  bool get isDark => _isDark;

  final darkTheme = ThemeData(
    fontFamily: 'SourceSerif',
    colorScheme: ColorScheme(
      primary: Colors.white,
      onPrimary: const Color.fromARGB(255, 28, 28, 28),
      secondary: Colors.white,
      onSecondary: const Color.fromARGB(255, 28, 28, 28),
      brightness: Brightness.dark,
      error: Colors.red,
      onError: Colors.white,
      surface: const Color.fromARGB(255, 28, 28, 28),
      onSurface: Colors.white,
    ),
    textTheme: const TextTheme(
      bodyLarge: TextStyle(fontSize: 16, height: 1.6, color: Colors.white),
      bodyMedium: TextStyle(fontSize: 14, height: 1.6, color: Colors.white),
      bodySmall: TextStyle(fontSize: 12, height: 1.6, color: Colors.white),
      headlineLarge: TextStyle(
          fontSize: 32,
          fontWeight: FontWeight.bold,
          color: Color.fromARGB(255, 207, 149, 33)),
      headlineMedium: TextStyle(
          fontSize: 28,
          fontWeight: FontWeight.bold,
          color: Color.fromARGB(255, 207, 149, 33)),
      headlineSmall: TextStyle(
          fontSize: 26,
          fontWeight: FontWeight.w600,
          color: Color.fromARGB(255, 207, 149, 33)),
    ),
    // Customize TextButton
    textButtonTheme: TextButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: Color.fromARGB(255, 207, 149, 33),
        foregroundColor: Color.fromARGB(255, 28, 28, 28),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: Color.fromARGB(255, 207, 149, 33),
        foregroundColor: Color.fromARGB(255, 28, 28, 28),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    outlinedButtonTheme: OutlinedButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: Color.fromARGB(255, 207, 149, 33),
        foregroundColor: Color.fromARGB(255, 28, 28, 28),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    dropdownMenuTheme: DropdownMenuThemeData(
      textStyle: TextStyle(fontSize: 16, color: Colors.white),
    ),
    dataTableTheme: DataTableThemeData(
      headingRowColor: WidgetStateProperty.all<Color>(
        const Color.fromARGB(255, 28, 28, 28), // Dark background for headers
      ),
      headingTextStyle: const TextStyle(
        color: Colors.white,
        fontWeight: FontWeight.bold,
        fontSize: 16,
      ),
      dataRowColor: WidgetStateProperty.resolveWith<Color?>(
        (states) {
          if (states.contains(WidgetState.hovered)) {
            return Colors.blue.withOpacity(0.1); // Change opacity when hovered
          }
          if (states.contains(WidgetState.selected)) {
            return Colors.blue.withOpacity(0.3); // Change opacity when selected
          }
          return null; // Default
        },
      ),
      dataTextStyle: const TextStyle(
        color: Color.fromARGB(255, 255, 255, 255),
        fontSize: 14,
      ),
      dividerThickness: 0, // Removes borders between rows
      columnSpacing: 16.0, // Spacing between columns
      horizontalMargin: 12.0, // Left and right margin for cells
      checkboxHorizontalMargin: 8.0, // Margin for checkboxes
    ),
  );

  final lightTheme = ThemeData(
    fontFamily: 'SourceSerif',
    colorScheme: ColorScheme(
      primary: const Color.fromARGB(255, 28, 28, 28),
      onPrimary: Colors.white,
      secondary: const Color.fromARGB(255, 28, 28, 28),
      onSecondary: Colors.white,
      brightness: Brightness.light,
      error: Colors.red,
      onError: Colors.white,
      surface: Colors.white,
      onSurface: const Color.fromARGB(255, 28, 28, 28),
    ),
    textTheme: const TextTheme(
      bodyLarge: TextStyle(
          fontSize: 16,
          height: 1.6,
          color: const Color.fromARGB(255, 28, 28, 28)),
      bodyMedium: TextStyle(
          fontSize: 14,
          height: 1.6,
          color: const Color.fromARGB(255, 28, 28, 28)),
      bodySmall: TextStyle(
          fontSize: 12,
          height: 1.6,
          color: const Color.fromARGB(255, 28, 28, 28)),
      headlineLarge: TextStyle(
          fontSize: 32,
          fontWeight: FontWeight.bold,
          color: const Color.fromARGB(255, 28, 28, 28)),
      headlineMedium: TextStyle(
          fontSize: 28,
          fontWeight: FontWeight.bold,
          color: const Color.fromARGB(255, 28, 28, 28)),
      headlineSmall: TextStyle(
          fontSize: 26,
          fontWeight: FontWeight.w600,
          color: const Color.fromARGB(255, 28, 28, 28)),
    ),
    // Customize TextButton
    textButtonTheme: TextButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: const Color.fromARGB(255, 28, 28, 28),
        foregroundColor: Colors.white,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: const Color.fromARGB(255, 28, 28, 28),
        foregroundColor: Colors.white,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    outlinedButtonTheme: OutlinedButtonThemeData(
      style: TextButton.styleFrom(
        backgroundColor: const Color.fromARGB(255, 28, 28, 28),
        foregroundColor: Colors.white,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      ),
    ),
    dropdownMenuTheme: DropdownMenuThemeData(
      textStyle: TextStyle(fontSize: 16, color: Colors.white),
    ),
    dataTableTheme: DataTableThemeData(
      headingRowColor: WidgetStateProperty.all<Color>(
        Colors.white, // Light background for headers
      ),
      headingTextStyle: const TextStyle(
        color: Color.fromARGB(255, 28, 28, 28),
        fontWeight: FontWeight.bold,
        fontSize: 16,
      ),
      dataRowColor: WidgetStateProperty.resolveWith<Color?>(
        (states) {
          if (states.contains(WidgetState.hovered)) {
            return Colors.blue.withOpacity(0.1); // Change opacity when hovered
          }
          if (states.contains(WidgetState.selected)) {
            return Colors.blue.withOpacity(0.3); // Change opacity when selected
          }
          return null; // Default
        },
      ),
      dataTextStyle: const TextStyle(
        color: const Color.fromARGB(255, 28, 28, 28),
        fontSize: 14,
      ),
      dividerThickness: 0, // Removes borders between rows
      columnSpacing: 16.0, // Spacing between columns
      horizontalMargin: 12.0, // Left and right margin for cells
      checkboxHorizontalMargin: 8.0, // Margin for checkboxes
    ),
  );

  ThemeData getTheme() {
    return _isDark ? darkTheme : lightTheme;
  }

  void toggleTheme() {
    _isDark = !_isDark;
    notifyListeners();
  }
}
