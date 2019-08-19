import React from 'react';
import ButtonWithIcon from '../../button_with_icon';

interface Props {
  searchText: string;
  handleSearchTextChange: (e: React.FormEvent<HTMLInputElement>) => void;
  handleSearchTextSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
}

const AppHeader: React.SFC<Props> = ({searchText, handleSearchTextChange, handleSearchTextSubmit}) => (
  <div className="flex justify-between items-center md:px-12 xs:px-3">
    <div className="md:hidden w-20 h-full">
      <ButtonWithIcon className="w-5" icon="search" />
    </div>

    <div className="flex-1 text-center">
      <h3 className="text-3xl font-bold text-indigo-700 leading-relaxed">phenopod</h3>
    </div>

    <div className="flex justify-around">
      <div className="md:block hidden flex-1 flex w-64 h-auto px-2 py-1">
        <form className="w-full" onSubmit={handleSearchTextSubmit}>
          <input className="w-full h-8 bg-gray-200 text-gray-800 px-2 py-1 rounded-lg" 
            type="text"
            placeholder="Search podcasts..."
            value={searchText}
            onChange={handleSearchTextChange}
          />
        </form>
      </div>

      <button className="w-20 h-8 flex-none border-2 border-orange-600 rounded-lg">
        <p className="w-full text-orange-700 text-sm font-semibold leading-loose text-center">SIGN IN</p>
      </button>
    </div>
  </div>
)

export default AppHeader