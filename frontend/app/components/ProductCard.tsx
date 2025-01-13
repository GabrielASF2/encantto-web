"use client";

import { useState } from 'react';
import { Product } from './product'; // Interface Product

interface ProductCardProps {
  product: Product;
}

export default function ProductCard({ product }: ProductCardProps) {
  const [selectedProduct, setSelectedProduct] = useState<number | null>(null);

  // Função para controlar a seleção
  const handleSelectProduct = (id: number) => {
    if (selectedProduct === id) {
      setSelectedProduct(null); // Desseleciona
    } else {
      setSelectedProduct(id); // Seleciona o card
    }
  };

  // Fechar ao clicar fora do card
  const handleClickOutside = (event: React.MouseEvent) => {
    const target = event.target as HTMLElement;
    if (target.closest('.card-content') === null) {
      setSelectedProduct(null); // Fecha se clicar fora
    }
  };

  return (
    <>
      {/* Container do background desfocado quando o card é selecionado */}
      {selectedProduct && (
        <div
          className="fixed inset-0 z-40 bg-black bg-opacity-50 backdrop-blur-md" // Desfoque de fundo
          onClick={handleClickOutside}
        ></div>
      )}

      <div
        className={`relative shadow-lg rounded-lg overflow-hidden bg-white transition-transform duration-300 ease-in-out cursor-pointer card-content flex flex-col ${
          selectedProduct === product.id ? 'z-50 scale-105' : ''
        }`}
        onClick={() => handleSelectProduct(product.id)} // Seleciona o card
        style={{
          maxWidth: selectedProduct === product.id ? 'auto' : 'auto',
          height: selectedProduct === product.id ? 'auto' : 'auto',
        }}
      >
        <img
          src={product.imageUrl}
          alt={product.title}
          className="object-cover w-full h-64"
        />
        <div className="p-4 flex flex-col flex-grow">
          <h2 className="text-xl font-bold text-gray-800">{product.title}</h2>
          <p className="mt-2 text-gray-600">{product.description}</p>
          <p className="mt-2 text-red-600 font-bold">
            R$ {product.price.toFixed(2)}
          </p>

          {/* Exibir customizações apenas quando o card estiver selecionado */}
          {selectedProduct === product.id && (
            <div className="mt-4 space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Nome para personalizar a capa:
                </label>
                <input
                  type="text"
                  className="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm"
                  onClick={(e) => e.stopPropagation()} // Evita fechar ao clicar no input
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Tipo de capa:
                </label>
                <select
                  className="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm"
                  onClick={(e) => e.stopPropagation()} // Evita fechar ao clicar no select
                >
                  <option value="Fosca">Fosca</option>
                  <option value="Holográfica">Holográfica</option>
                  <option value="Tradicional">Tradicional</option>
                </select>
              </div>
            </div>
          )}

          {/* Botão "Comprar" sempre visível */}
          <button
            className="mt-4 bg-blue-600 text-white px-4 py-2 rounded-md shadow hover:bg-blue-700 w-full"
            onClick={(e) => e.stopPropagation()} // Evita fechar ao clicar no botão
          >
            Comprar
          </button>
        </div>

        {/* Botão para fechar o card */}
        {selectedProduct === product.id && (
          <button
            onClick={(e) => {
              e.stopPropagation(); // Evita fechar diretamente pelo clique no modal
              setSelectedProduct(null); // Fecha o card manualmente
            }}
          >
          </button>
        )}
      </div>
    </>
  );
}
