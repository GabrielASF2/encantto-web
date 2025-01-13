export interface Product {
    id: number;
    title: string;
    price: number;
    description: string;
    imageUrl: string;
    customization: boolean;
  }
  
  export const products: Product[] = [
    {
      id: 1,
      title: "Caderno Personalizado",
      price: 44.90,
      description: "Caderno de encadernação artesanal, perfeito para anotações e diário.",
      imageUrl: "/path-to-image1.jpg", // Altere para a URL da sua imagem
      customization: true,
    },
    {
      id: 2,
      title: "Agenda 2024",
      price: 54.90,
      description: "Agenda personalizada para planejar seu ano com estilo.",
      imageUrl: "/path-to-image2.jpg",
      customization: true,
    },
    {
        id: 3,
        title: "Caderno de Receitas",
        price: 39.90,
        description: "Caderno feito à mão para guardar suas receitas favoritas.",
        imageUrl: "/path-to-image3.jpg",
        customization: true,
    },
    {
        id: 4,
        title: "Capa de Livro Personalizada",
        price: 29.90,
        description: "Capa de livro feita sob medida para seus livros favoritos.",
        imageUrl: "/path-to-image4.jpg",
        customization: true,
    },
    {
        id: 5,
        title: "Caderno de Viagem",
        price: 49.90,
        description: "Capture suas memórias de viagem com nosso caderno especial.",
        imageUrl: "/path-to-image5.jpg",
        customization: true,
    },
    {
        id: 6,
        title: "Caderno de Anotações",
        price: 34.90,
        description: "Perfeito para estudantes ou profissionais, com papel de alta qualidade.",
        imageUrl: "/path-to-image6.jpg",
        customization: false,
    }
    // Outros produtos...
  ];
  