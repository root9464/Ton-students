type ServiceInfo = {
  title: string;
  content: string;
};

type Service = {
  id: string;
  price: number;
  infos: ServiceInfo[];
  buttonText?: string; // Поле опционально, так как не во всех записях оно есть
};

type ApiResponse = {
  data: Service[];
  message: string;
  status: string;
};

const response: ApiResponse = {
  message: 'Services fetched successfully',
  data: [
    {
      id: '105380b2-a06c-48db-9454-0b5c084af76a',
      price: 0.2,
      infos: [
        {
          title: 'Service Title 1',
          content: 'Detailed content for service 1.',
        },
      ],
      buttonText: 'fffffff',
    },
  ],
  status: 'success',
};

const response2: ApiResponse = {
  data: [
    {
      id: '105380b2-a06c-48db-9454-0b5c084af76a',
      price: 0.2,
      infos: [
        {
          title: 'Service Title 1',
          content: 'Detailed content for service 1.',
        },
      ],
      buttonText: 'fffffff',
    },
  ],
  message: 'Services fetched successfully',
  status: 'success',
};

console.log(response2 === response);
