import ProductCard from '../components/ProductCard';
import { products } from '../components/product'; // Importa os produtos

export default function Catalog() {
  return (
    <div className="w-full flex flex-col items-center py-8">
      <h1 className="text-5xl font-bold text-[#8DBBB8] mb-6">Nosso Catálogo</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 max-w-6xl w-full px-4">
        {products.map((product) => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
}