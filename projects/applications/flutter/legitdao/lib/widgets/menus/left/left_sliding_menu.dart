import 'package:flutter/material.dart';
import 'package:easy_localization/easy_localization.dart';
import '../../networks/network_manager_interface.dart';
import '../../visuals/buttons/hover_link.dart';

class LeftSlidingMenu extends StatefulWidget {
  final bool isVisible;
  final Function() onClose;
  final NetworkManager networkManager;

  const LeftSlidingMenu({
    super.key,
    required this.isVisible,
    required this.onClose,
    required this.networkManager,
  });

  @override
  _LeftSlidingMenuState createState() => _LeftSlidingMenuState();
}

class _LeftSlidingMenuState extends State<LeftSlidingMenu> {
  late bool _isVisible;

  @override
  void initState() {
    super.initState();
    _isVisible = widget.isVisible;
  }

  @override
  void didUpdateWidget(covariant LeftSlidingMenu oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.isVisible != widget.isVisible) {
      setState(() {
        _isVisible = widget.isVisible;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    double width = MediaQuery.of(context).size.width;
    if (width < 500.0) {
      width = 500.0;
    }

    if (width > 800.0) {
      width = 800.0;
    }

    width = width * 0.6;

    return AnimatedPositioned(
      duration: const Duration(milliseconds: 300),
      curve: Curves.easeInOut,
      top: 0,
      bottom: 0,
      left: _isVisible ? 0 : (width * -1.0),
      width: width,
      child: KeyedSubtree(
        key: ValueKey(_isVisible),
        child: Container(
          decoration: BoxDecoration(
            border: Border(
              right: BorderSide(
                width: 2.0,
              ),
            ),
          ),
          child: Material(
            elevation: 4,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                // X Close Button
                Container(
                  padding: const EdgeInsets.all(16.0), // Optional padding
                  child: Align(
                    alignment: Alignment.topRight, // Align to the top-right
                    child: Container(
                      decoration: BoxDecoration(
                        color: Colors.grey[200], // Background color
                        borderRadius:
                            BorderRadius.circular(8.0), // Rounded corners
                      ),
                      child: IconButton(
                        icon: const Icon(Icons.close),
                        color: Colors.black, // Icon color
                        onPressed: () {
                          widget.onClose();
                        },
                      ),
                    ),
                  ),
                ),

                // Menu Items
                Expanded(
                  child: SingleChildScrollView(
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        // Home
                        Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: HoverLink(
                            text: "menu_home".tr(),
                            isInMenu: true,
                            onTap: () {
                              Navigator.pushNamed(context, '/')
                                  .then((_) => widget.onClose());
                            },
                          ),
                        ),

                        // Tokens
                        Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: HoverLink(
                            text: "menu_tokens".tr(),
                            isInMenu: true,
                            onTap: () {
                              Navigator.pushNamed(context, '/tokens')
                                  .then((_) => widget.onClose());
                            },
                          ),
                        ),

                        // About
                        Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: HoverLink(
                            text: "menu_about".tr(),
                            isInMenu: true,
                            onTap: () {
                              Navigator.pushNamed(context, '/about')
                                  .then((_) => widget.onClose());
                            },
                          ),
                        ),

                        // Contact
                        Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: HoverLink(
                            text: "menu_contact".tr(),
                            isInMenu: true,
                            onTap: () {
                              Navigator.pushNamed(context, '/contact')
                                  .then((_) => widget.onClose());
                            },
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
