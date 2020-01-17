import classnames from 'classnames'
import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'

interface Props {
  searchText: string
  handleSearchTextChange: (e: React.FormEvent<HTMLInputElement>) => void
  handleSearchTextSubmit: (e: React.FormEvent<HTMLFormElement>) => void
}

const SearchBar: React.SFC<Props> = (props) => {
  const { searchText, handleSearchTextChange, handleSearchTextSubmit } = props

  return (
    <form
      className="relative flex items-center px-2 py-1"
      style={{ width: '38rem' }}
      onSubmit={handleSearchTextSubmit}
    >
      <ButtonWithIcon
        className="absolute inset-y-0 right-0 w-5 mr-4 text-gray-600"
        icon="search"
        type="submit"
      />
      <input
        className={classnames(
          'w-full h-9 pl-4 pr-8 py-1 text-gray-900 tracking-wide placeholder-gray-700 border border-gray-400 rounded-full',
          'appearance-none focus:outline-none',
        )}
        type="text"
        placeholder="Search for topic, person, podcast, episode..."
        value={searchText}
        onChange={handleSearchTextChange}
      />
    </form>
  )
}

export default SearchBar
