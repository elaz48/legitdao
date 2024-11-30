import 'package:flutter/material.dart';
import 'package:easy_localization/easy_localization.dart';
import '../widgets/referrals/referrals.dart';

class ReferralsPage extends StatefulWidget {
  final String walletAddress;

  ReferralsPage({
    super.key,
    required this.walletAddress,
  });

  @override
  _ReferralsState createState() => _ReferralsState();
}

class _ReferralsState extends State<ReferralsPage>
    with TickerProviderStateMixin {
  String walletAddress = "";
  @override
  void initState() {
    super.initState();
    walletAddress = widget.walletAddress;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("ReferralsPage Tree").tr(),
      ),
      body: Container(
        height: 80,
        color: Colors.blue,
        child: Center(
          child: Referrals(walletAddress: walletAddress),
        ),
      ),
    );
  }
}
