/* eslint-disable @typescript-eslint/no-unused-expressions */
import { $debug, backButton, createPostEvent, initData, init as initSDK, miniApp, themeParams, viewport } from '@telegram-apps/sdk-react';

/**
 * Initializes the application and configures its dependencies.
 */
export function init(debug: boolean): void {
  // Set @telegram-apps/sdk-react debug mode.
  $debug.set(debug);
  const postEvent = createPostEvent('8.0');
  postEvent('web_app_request_fullscreen');

  // Initialize special event handlers for Telegram Desktop, Android, iOS, etc.
  // Also, configure the package.
  initSDK();

  // Mount all components used in the project.
  backButton.isSupported() && backButton.mount();
  miniApp.mount();
  viewport.mount();
  themeParams.mount();

  // settings
  viewport.expand();
  initData.restore();
  backButton.hide();

  // Add Eruda if needed.
  debug && import('eruda').then((lib) => lib.default.init()).catch(console.error);
}
