'use client';
import { useBackButton } from '@/shared/hooks/useBackButton';
import { Button } from '@nextui-org/button';

export default function ProfilePage() {
  useBackButton(true);

  return (
    <div className='h-contentFlow w-full'>
      <Button className='h-[50px] w-[100px] bg-uiLightBLueGradient text-uiDeepLightBlue'>fffff</Button>
    </div>
  );
}
