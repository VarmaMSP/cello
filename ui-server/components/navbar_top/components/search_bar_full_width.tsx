import React from 'react';
import classnames from 'classnames';
import ButtonWithIcon from '../../button_with_icon';

interface Props {
  searchText: string;
  handleCollapse: () => void;
  handleSearchTextChange: (e: React.FormEvent<HTMLInputElement>) => void;
  handleSearchTextSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
}

const SearchBarFullWidth: React.SFC<Props> = (props) => {
  const {
    searchText,
    handleCollapse,
    handleSearchTextChange,
    handleSearchTextSubmit,
  } = props

  return <form className="relative flex items-center h-full w-full px-2 py-1" onSubmit={handleSearchTextSubmit}>
    <ButtonWithIcon 
      className="absolute inset-y-0 left-0 w-5 ml-4 text-gray-700" 
      icon="arrow-left"
      onClick={handleCollapse}
    />
    <ButtonWithIcon
      className="absolute inset-y-0 right-0 w-5 mr-4 text-gray-700"
      icon="search"
      type="submit"
    />
    <input
      className={classnames(
        "w-full h-full px-10 py-1 bg-gray-200 text-gray-800 placeholder-gray-600 border border-transparent rounded-lg",
        "appearance-none focus:outline-none focus:bg-white focus:border-gray-400"
      )}
      type="text"
      placeholder="Search podcasts"
      value={searchText}
      onChange={handleSearchTextChange}
    />
  </form>
}

export default SearchBarFullWidth
