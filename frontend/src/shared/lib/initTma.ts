/* eslint-disable @typescript-eslint/no-unused-expressions */
import { $debug, backButton, initData, init as initSDK, miniApp, themeParams } from '@telegram-apps/sdk-react';
/**
 * Initializes the application and configures its dependencies.
 */
export function init(debug: boolean): void {
  // Set @telegram-apps/sdk-react debug mode.
  $debug.set(debug);

  // Initialize special event handlers for Telegram Desktop, Android, iOS, etc.
  // Also, configure the package.
  initSDK();

  // Mount all components used in the project.
  backButton.isSupported() && backButton.mount();
  miniApp.mount();
  themeParams.mount();
  initData.restore();

  // Add Eruda if needed.
  debug && import('eruda').then((lib) => lib.default.init()).catch(console.error);
}
