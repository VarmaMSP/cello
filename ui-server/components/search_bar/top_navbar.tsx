import classnames from 'classnames'
import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import withProps, { SearchBarProps } from './with_props'

const SearchBar: React.FC<SearchBarProps> = ({
  searchText,
  handleTextChange,
  handleTextSubmit,
  collapseSearchBar,
}) => {
  return (
    <form
      className="relative flex items-center h-full w-full px-2 py-1"
      onSubmit={handleTextSubmit}
    >
      <ButtonWithIcon
        className="absolute inset-y-0 left-0 w-5 ml-5 text-gray-700"
        icon="arrow-left"
        onClick={collapseSearchBar}
      />
      <ButtonWithIcon
        className="absolute inset-y-0 right-0 w-5 mr-5 text-gray-700"
        icon="search"
        type="submit"
      />
      <input
        className={classnames(
          'w-full h-full px-12 py-1 bg-white placeholder-gray-600 border border-gray-400 rounded-lg',
          'appearance-none focus:outline-none focus:border-blue-500',
        )}
        type="text"
        placeholder="Search podcasts"
        value={searchText}
        onChange={handleTextChange}
      />
    </form>
  )
}

export default withProps(SearchBar)
