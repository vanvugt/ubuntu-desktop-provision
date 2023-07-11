import 'ubuntu_provision_localizations.dart';

/// The translations for Italian (`it`).
class UbuntuProvisionLocalizationsIt extends UbuntuProvisionLocalizations {
  UbuntuProvisionLocalizationsIt([String locale = 'it']) : super(locale);

  @override
  String get timezonePageTitle => 'Select your timezone';

  @override
  String get timezoneLocationLabel => 'Location';

  @override
  String get timezoneTimezoneLabel => 'Fuso orario';

  @override
  String get keyboardTitle => 'Disposizione della tastiera';

  @override
  String get keyboardHeader => 'Scegliere la disposizione della tastiera:';

  @override
  String get keyboardTestHint => 'Digitare qui per provare la tastiera';

  @override
  String get keyboardDetectTitle => 'Rileva disposizione tastiera';

  @override
  String get keyboardDetectButton => 'Detect';

  @override
  String get keyboardVariantLabel => 'Keyboard variant:';

  @override
  String get keyboardPressKeyLabel => 'Prego premere uno dei seguenti tasti:';

  @override
  String get keyboardKeyPresentLabel => 'Il seguente tasto è presente sulla propria tastiera?';

  @override
  String get themePageTitle => 'Scegli il aspetto';

  @override
  String get themePageHeader => 'You can always change this later in the appearance settings.';

  @override
  String get themeDark => 'Dark';

  @override
  String get themeLight => 'Light';

  @override
  String localePageTitle(Object DISTRO) {
    return 'Welcome to $DISTRO';
  }

  @override
  String get localeHeader => 'Scegli la tua lingua:';

  @override
  String get identityPageTitle => 'Informazioni personali';

  @override
  String get identityAutoLogin => 'Accedere automaticamente';

  @override
  String get identityRequirePassword => 'Require my password to log in';

  @override
  String get identityRealNameLabel => 'Your name';

  @override
  String get identityRealNameRequired => 'A name is required';

  @override
  String get identityRealNameTooLong => 'That name is too long.';

  @override
  String get identityHostnameLabel => 'Il nome del computer';

  @override
  String get identityHostnameInfo => 'The name it uses when it talks to other computers.';

  @override
  String get identityHostnameRequired => 'A computer name is required';

  @override
  String get identityHostnameTooLong => 'That computer name is too long.';

  @override
  String get identityInvalidHostname => 'The computer name is invalid';

  @override
  String get identityUsernameLabel => 'Scegli uno username';

  @override
  String get identityUsernameRequired => 'A username is required';

  @override
  String get identityInvalidUsername => 'The username is invalid';

  @override
  String get identityUsernameInUse => 'That user name already exists.';

  @override
  String get identityUsernameSystemReserved => 'That name is reserved for system usage.';

  @override
  String get identityUsernameTooLong => 'That name is too long.';

  @override
  String get identityUsernameInvalidChars => 'That name contains invalid characters.';

  @override
  String get identityPasswordLabel => 'Scegli una password';

  @override
  String get identityPasswordRequired => 'A password is required';

  @override
  String get identityConfirmPasswordLabel => 'Conferma la password';

  @override
  String get identityPasswordMismatch => 'The passwords do not match';

  @override
  String get identityPasswordShow => 'Mostra';

  @override
  String get identityPasswordHide => 'Nascondi';

  @override
  String get identityActiveDirectoryOption => 'Use Active Directory';

  @override
  String get identityActiveDirectoryInfo => 'You\'ll enter domain and other details in the next step.';
}
