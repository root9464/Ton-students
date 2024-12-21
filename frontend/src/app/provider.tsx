'use client';
import { useClientOnce } from '@/shared/hooks/useClientOnce';
import '@/shared/hooks/useTelegramMock';
import { init } from '@/shared/lib/initTma';
import { NextUIProvider } from '@nextui-org/system';
import { ThemeProvider as NextThemesProvider } from 'next-themes';
import { ReactNode } from 'react';

import '@/shared/util/mock';

export const Provider = ({ children }: Readonly<{ children: ReactNode }>) => {
  useClientOnce(async () => {
    init(true);
  });

  return (
    <NextUIProvider className='h-full w-full'>
      <NextThemesProvider attribute='class' defaultTheme='light'>
        {children}
      </NextThemesProvider>
    </NextUIProvider>
  );
};
