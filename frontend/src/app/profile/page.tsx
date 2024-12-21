'use client';
import { Button } from '@nextui-org/button';
import { backButton } from '@telegram-apps/sdk';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function ProfilePage() {
  const router = useRouter();

  useEffect(() => {
    backButton.mount();
    backButton.show();
    return backButton.onClick(() => {
      router.back();
    });
  }, [router]);
  return (
    <div className='h-contentFlow w-full'>
      <Button className='h-[50px] w-[100px] bg-uiLightBLueGradient text-uiDeepLightBlue'>fffff</Button>
    </div>
  );
}
