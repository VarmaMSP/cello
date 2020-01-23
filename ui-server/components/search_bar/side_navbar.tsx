import classnames from 'classnames'
import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import withProps, { SearchBarProps } from './with_props'

const SearchBar: React.FC<SearchBarProps> = ({
  searchText,
  handleTextChange,
  handleTextSubmit,
}) => {
  return (
    <form
      className="relative flex items-center px-2 py-1"
      onSubmit={handleTextSubmit}
    >
      <ButtonWithIcon
        className="absolute inset-y-0 right-0 w-4 h-auto mr-4 text-gray-700"
        icon="search"
        type="submit"
      />
      <input
        className={classnames(
          'w-full h-8 pl-4 pr-6 py-1 text-gray-900 tracking-wide placeholder-gray-700 bg-white border border-gray-500 rounded-full',
          'appearance-none focus:outline-none focus:border-2 focus:border-blue-500',
        )}
        type="text"
        placeholder="Search..."
        value={searchText}
        onChange={handleTextChange}
      />
    </form>
  )
}

export default withProps(SearchBar)
