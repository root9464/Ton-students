import { mockTelegramEnv } from '@telegram-apps/bridge';

export const launchParams = () =>
  mockTelegramEnv({
    themeParams: {
      accentTextColor: '#6ab2f2',
      bgColor: '#17212b',
      buttonColor: '#5288c1',
      buttonTextColor: '#ffffff',
      destructiveTextColor: '#ec3942',
      headerBgColor: '#17212b',
      hintColor: '#708499',
      linkColor: '#6ab3f3',
      secondaryBgColor: '#232e3c',
      sectionBgColor: '#17212b',
      sectionHeaderTextColor: '#6ab3f3',
      subtitleTextColor: '#708499',
      textColor: '#f5f5f5',
    },
    initData: {
      user: {
        id: 99281932,
        firstName: 'Andrew',
        lastName: 'Rogue',
        username: 'rogue',
        languageCode: 'en',
        isPremium: true,
        allowsWriteToPm: true,
      },
      signature: 'ff',
      hash: '89d6079ad6762351f38c6dbbc41bb53048019256a9443988af7a48bcad16ba31',
      authDate: new Date(1716922846000),
      startParam: 'debug',
      chatType: 'sender',
      chatInstance: '8428209589180549439',
    },
    initDataRaw: new URLSearchParams([
      [
        'user',
        JSON.stringify({
          id: 99281932,
          first_name: 'Andrew',
          last_name: 'Rogue',
          username: 'rogue',
          language_code: 'en',
          is_premium: true,
          allows_write_to_pm: true,
          photo_url: 'https://t4.ftcdn.net/jpg/00/88/86/83/360_F_88868306_YMrcsGUTR0DSEQ44vvsWDqROVt1Kk4V9.jpg',
        }),
      ],
      ['hash', '89d6079ad6762351f38c6dbbc41bb53048019256a9443988af7a48bcad16ba31'],
      ['auth_date', '1716922846'],
      ['start_param', 'debug'],
      ['chat_type', 'sender'],
      ['chat_instance', '8428209589180549439'],
    ]).toString(),
    version: '7.2',
    platform: 'tdesktop',
  });
