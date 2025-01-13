import { FaInstagram, FaWhatsapp } from 'react-icons/fa';
import Catalog from './team/page';

export default function Home() {
  return (
    <div className="w-auto flex flex-col justify-center items-center py-8"> {/* Adicionando padding */}
      <div className="shadow-lg rounded-lg overflow-hidden max-w-4xl w-full flex flex-col lg:flex-row">
        {/* Imagem */}
        <div className="relative w-full lg:w-1/2 flex flex-col justify-center">
          <div className="relative m-4">
            <img
              src="/logo.svg"
              alt="Handmade Craft"
              className="object-cover w-full h-full rounded-md"
            />
            
          </div>
        </div>

        {/* Texto e botão */}
        <div className="p-8 w-full lg:w-1/2 flex flex-col"> {/* Adicionando flex-col para o texto */}
          <h2 className="text-4xl font-bold text-[#8DBBB8]">by Nathali Franco</h2>
          <p className="mt-4 text-gray-600">
            Bem-vindo à [Nome da Loja], onde a arte da encadernação se encontra com a criatividade! Especializados em produtos de papelaria de alta qualidade, oferecemos uma vasta gama de agendas customizadas, cadernos e artigos de papelaria que atendem a todos os estilos e necessidades.
            <br /><br />
            Venha nos visitar e descubra como podemos tornar sua experiência de papelaria única e especial. Na [Nome da Loja], a sua criatividade é a nossa inspiração!
          </p>
          <button className="mt-6 bg-[#EAAEB2] text-white px-6 py-2 rounded-md shadow hover:bg-[#F5B4A5]">
            <a href="/team"> {/* Adicionando link para a página de produtos */}
           Confira alguns dos nossos produtos!
            </a>
          </button>

          {/* Ícones de redes sociais */}
          <div className="flex space-x-4 mt-6 justify-center">
            <a href="https://wa.me/554898523303" className="text-gray-600 hover:text-green-500">
              <FaWhatsapp size={24} />
            </a>
            <a href="https://www.instagram.com/encantto_papelaria/" className="text-gray-600 hover:text-pink-600">
              <FaInstagram size={24} />
            </a>
          </div>
        </div>
      </div>
      <Catalog>
      </Catalog>
    </div>
  );
}
