import classnames from 'classnames'
import ButtonWithIcon from 'components/button_with_icon'
import SearchSuggestions from 'components/search_suggestions'
import usePopper from 'hooks/usePopper'
import React from 'react'
import { Portal } from 'react-portal'
import { stopEventPropagation } from 'utils/dom'
import withProps, { SearchBarProps } from './with_props'

const SearchBar: React.FC<SearchBarProps> = ({
  searchText,
  handleTextChange,
  handleTextSubmit,
  showSuggestions,
  setShowSuggestions,
}) => {
  const [reference, popper] = usePopper(
    {
      placement: 'bottom-start',
      modifiers: [
        {
          name: 'offset',
          options: {
            offset: [0, 5],
          },
        },
      ],
      strategy: 'fixed',
    },
    () => setShowSuggestions(false),
  )

  return (
    <>
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
            'w-full h-8 pl-4 pr-6 py-1 text-gray-900 tracking-wide placeholder-gray-700 bg-white border border-gray-500 rounded-lg',
            'appearance-none focus:outline-none focus:border-2 focus:border-blue-500',
          )}
          type="text"
          placeholder="Search..."
          value={searchText}
          onChange={handleTextChange}
          ref={reference.ref}
          onFocus={() => setShowSuggestions(true)}
          onPointerDown={stopEventPropagation}
          onMouseDown={stopEventPropagation}
          onTouchStart={stopEventPropagation}
        />
      </form>

      {showSuggestions && (
        <Portal>
          <div ref={popper.ref} style={popper.styles}>
            <SearchSuggestions />
          </div>
        </Portal>
      )}
    </>
  )
}

export default withProps(SearchBar)
