import { backButton } from '@telegram-apps/sdk-react';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export const useBackButton = (visible: boolean) => {
  const router = useRouter();

  useEffect(() => {
    if (visible) {
      backButton.show();
    } else {
      backButton.hide();
    }
  }, [router, visible]);

  useEffect(() => {
    return backButton.onClick(() => {
      router.back();
    });
  }, [router]);

  return null;
};
