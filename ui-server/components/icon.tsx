import React from 'react'

export type Icon =
  | 'add-outline'
  | 'arrow-left'
  | 'arrow-right'
  | 'cheveron-up'
  | 'close'
  | 'facebook-color'
  | 'fast-forward'
  | 'fast-rewind'
  | 'feed'
  | 'google-color'
  | 'heart'
  | 'history'
  | 'home'
  | 'logo-lg'
  | 'logo-md'
  | 'minus-outline'
  | 'pause'
  | 'phenopod'
  | 'play'
  | 'play-outline'
  | 'search'
  | 'twitter-color'
  | 'user-solid-circle'
  | 'volume'
  | 'walk'

export const iconMap: { [key in Icon]: React.SFC<{ className: string }> } = {
  'add-outline': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M11 9h4v2h-4v4H9v-4H5V9h4V5h2v4zm-1 11a10 10 0 110-20 10 10 0 010 20zm0-2a8 8 0 100-16 8 8 0 000 16z" />
    </svg>
  ),

  'arrow-left': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M3.828 9l6.071-6.071-1.414-1.414L0 10l.707.707 7.778 7.778 1.414-1.414L3.828 11H20V9H3.828z" />
    </svg>
  ),

  'arrow-right': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M16.172 9l-6.071-6.071 1.414-1.414L20 10l-.707.707-7.778 7.778-1.414-1.414L16.172 11H0V9z" />
    </svg>
  ),

  'cheveron-up': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10.707 7.05L10 6.343 4.343 12l1.414 1.414L10 9.172l4.243 4.242L15.657 12z" />
    </svg>
  ),

  close: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 8.586L2.929 1.515 1.515 2.929 8.586 10l-7.071 7.071 1.414 1.414L10 11.414l7.071 7.071 1.414-1.414L11.414 10l7.071-7.071-1.414-1.414L10 8.586z" />
    </svg>
  ),

  'facebook-color': (props) => (
    <svg viewBox="0 0 408.788 408.788" {...props}>
      <path
        d="M353.701 0H55.087C24.665 0 .002 24.662.002 55.085v298.616c0 30.423 24.662 55.085 55.085 55.085h147.275l.251-146.078h-37.951a8.954 8.954 0 01-8.954-8.92l-.182-47.087a8.955 8.955 0 018.955-8.989h37.882v-45.498c0-52.8 32.247-81.55 79.348-81.55h38.65a8.955 8.955 0 018.955 8.955v39.704a8.955 8.955 0 01-8.95 8.955l-23.719.011c-25.615 0-30.575 12.172-30.575 30.035v39.389h56.285c5.363 0 9.524 4.683 8.892 10.009l-5.581 47.087a8.955 8.955 0 01-8.892 7.901h-50.453l-.251 146.078h87.631c30.422 0 55.084-24.662 55.084-55.084V55.085C408.786 24.662 384.124 0 353.701 0z"
        fill="#475993"
      />
    </svg>
  ),

  'fast-forward': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M1 5l9 5-9 5V5zm9 0l9 5-9 5V5z" />
    </svg>
  ),

  'fast-rewind': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M19 5v10l-9-5 9-5zm-9 0v10l-9-5 9-5z" />
    </svg>
  ),

  feed: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M6 6V2c0-1.1.9-2 2-2h10a2 2 0 012 2v10a2 2 0 01-2 2h-4v4a2 2 0 01-2 2H2a2 2 0 01-2-2V8c0-1.1.9-2 2-2h4zm2 0h4a2 2 0 012 2v4h4V2H8v4zM2 8v10h10V8H2z" />
    </svg>
  ),

  'google-color': (props) => (
    <svg viewBox="0 0 512 512" {...props}>
      <path
        d="M113.47 309.408L95.648 375.94l-65.139 1.378C11.042 341.211 0 299.9 0 256c0-42.451 10.324-82.483 28.624-117.732h.014L86.63 148.9l25.404 57.644c-5.317 15.501-8.215 32.141-8.215 49.456.002 18.792 3.406 36.797 9.651 53.408z"
        fill="#fbbb00"
      />
      <path
        d="M507.527 208.176C510.467 223.662 512 239.655 512 256c0 18.328-1.927 36.206-5.598 53.451-12.462 58.683-45.025 109.925-90.134 146.187l-.014-.014-73.044-3.727-10.338-64.535c29.932-17.554 53.324-45.025 65.646-77.911h-136.89V208.176h245.899z"
        fill="#518ef8"
      />
      <path
        d="M416.253 455.624l.014.014C372.396 490.901 316.666 512 256 512c-97.491 0-182.252-54.491-225.491-134.681l82.961-67.91c21.619 57.698 77.278 98.771 142.53 98.771 28.047 0 54.323-7.582 76.87-20.818l83.383 68.262z"
        fill="#28b446"
      />
      <path
        d="M419.404 58.936l-82.933 67.896C313.136 112.246 285.552 103.82 256 103.82c-66.729 0-123.429 42.957-143.965 102.724l-83.397-68.276h-.014C71.23 56.123 157.06 0 256 0c62.115 0 119.068 22.126 163.404 58.936z"
        fill="#f14336"
      />
    </svg>
  ),

  heart: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 3.22l-.61-.6a5.5 5.5 0 00-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 00-7.78-7.77l-.61.61z" />
    </svg>
  ),

  history: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M3 18a7 7 0 014-6.33V8.33A7 7 0 013 2H1V0h18v2h-2a7 7 0 01-4 6.33v3.34A7 7 0 0117 18h2v2H1v-2h2zM5 2a5 5 0 004 4.9V10h2V6.9A5 5 0 0015 2H5z" />
    </svg>
  ),

  home: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M8 20H3V10H0L10 0l10 10h-3v10h-5v-6H8v6z" />
    </svg>
  ),

  'logo-lg': (props) => (
    <svg width={167} height={37} fill="none" {...props}>
      <path
        d="M11.605 28.352c-2.308 0-4.119-.838-5.431-2.514h-.281c.187 1.64.28 2.59.28 2.848v7.962H.814v-28.3h4.359l.756 2.548h.246C7.428 8.951 9.285 7.98 11.746 7.98c2.32 0 4.137.896 5.45 2.689 1.312 1.793 1.968 4.283 1.968 7.47 0 2.098-.31 3.92-.932 5.467-.609 1.547-1.482 2.725-2.619 3.534-1.136.808-2.472 1.213-4.008 1.213zm-1.582-16.084c-1.324 0-2.29.41-2.9 1.23-.61.809-.926 2.15-.95 4.025v.58c0 2.11.311 3.622.932 4.536.633.914 1.63 1.37 2.989 1.37 2.402 0 3.603-1.98 3.603-5.94 0-1.934-.299-3.381-.896-4.342-.586-.973-1.512-1.46-2.778-1.46zM40.294 28h-5.361V16.521c0-2.836-1.055-4.253-3.164-4.253-1.5 0-2.584.51-3.252 1.529-.668 1.02-1.002 2.672-1.002 4.957V28h-5.361V.648h5.361v5.573c0 .433-.04 1.453-.123 3.058l-.123 1.582h.281c1.195-1.922 3.094-2.882 5.695-2.882 2.31 0 4.061.62 5.256 1.863 1.196 1.242 1.793 3.023 1.793 5.343V28zm12.219-16.207c-1.137 0-2.028.363-2.672 1.09-.645.715-1.014 1.734-1.108 3.058h7.524c-.024-1.324-.37-2.343-1.037-3.058-.668-.727-1.57-1.09-2.707-1.09zm.755 16.559c-3.164 0-5.636-.873-7.418-2.62-1.78-1.746-2.671-4.218-2.671-7.418 0-3.293.82-5.835 2.46-7.629 1.653-1.804 3.932-2.706 6.838-2.706 2.778 0 4.94.79 6.487 2.373 1.547 1.582 2.32 3.767 2.32 6.556v2.602H48.61c.059 1.523.51 2.713 1.354 3.568.843.856 2.027 1.283 3.55 1.283a14.74 14.74 0 003.358-.369c1.055-.246 2.156-.639 3.305-1.178v4.149c-.938.469-1.94.814-3.006 1.037-1.067.234-2.367.352-3.903.352zM82.344 28h-5.361V16.521c0-1.418-.252-2.478-.756-3.181-.504-.715-1.307-1.072-2.408-1.072-1.5 0-2.584.503-3.252 1.511-.668.996-1.002 2.655-1.002 4.975V28h-5.362V8.348H68.3l.72 2.513h.3a5.619 5.619 0 012.46-2.144c1.055-.492 2.25-.738 3.587-.738 2.285 0 4.02.62 5.203 1.863 1.183 1.23 1.775 3.011 1.775 5.343V28zm8.351-9.861c0 1.945.317 3.416.95 4.412.644.996 1.687 1.494 3.128 1.494 1.43 0 2.455-.492 3.076-1.477.633-.996.95-2.472.95-4.43 0-1.945-.317-3.404-.95-4.376-.632-.973-1.67-1.46-3.111-1.46-1.43 0-2.46.487-3.094 1.46-.633.96-.949 2.42-.949 4.377zm13.588 0c0 3.199-.844 5.7-2.531 7.506-1.688 1.804-4.037 2.707-7.05 2.707-1.886 0-3.55-.41-4.991-1.23-1.442-.833-2.55-2.022-3.323-3.57-.773-1.546-1.16-3.35-1.16-5.413 0-3.211.838-5.707 2.514-7.489 1.676-1.78 4.031-2.671 7.066-2.671 1.887 0 3.551.41 4.992 1.23 1.442.82 2.549 1.998 3.323 3.533.773 1.535 1.16 3.334 1.16 5.397zm13.783 10.213c-2.309 0-4.12-.838-5.432-2.514h-.281c.187 1.64.281 2.59.281 2.848v7.962h-5.361v-28.3h4.359l.756 2.548h.246c1.254-1.945 3.111-2.917 5.572-2.917 2.321 0 4.137.896 5.449 2.689 1.313 1.793 1.969 4.283 1.969 7.47 0 2.098-.31 3.92-.931 5.467-.61 1.547-1.483 2.725-2.62 3.534-1.136.808-2.472 1.213-4.007 1.213zm-1.582-16.084c-1.325 0-2.291.41-2.901 1.23-.609.809-.926 2.15-.949 4.025v.58c0 2.11.311 3.622.932 4.536.632.914 1.629 1.37 2.988 1.37 2.402 0 3.603-1.98 3.603-5.94 0-1.934-.298-3.381-.896-4.342-.586-.973-1.512-1.46-2.777-1.46zm16.401 5.87c0 1.946.317 3.417.95 4.413.644.996 1.687 1.494 3.129 1.494 1.429 0 2.455-.492 3.076-1.477.633-.996.949-2.472.949-4.43 0-1.945-.316-3.404-.949-4.376-.633-.973-1.67-1.46-3.112-1.46-1.429 0-2.461.487-3.093 1.46-.633.96-.95 2.42-.95 4.377zm13.588 0c0 3.2-.843 5.702-2.531 7.506-1.687 1.805-4.037 2.708-7.049 2.708-1.887 0-3.551-.41-4.992-1.23-1.441-.833-2.549-2.022-3.322-3.57-.774-1.546-1.16-3.35-1.16-5.413 0-3.211.837-5.707 2.513-7.489 1.676-1.78 4.032-2.671 7.067-2.671 1.886 0 3.55.41 4.992 1.23 1.441.82 2.549 1.998 3.322 3.533.774 1.535 1.16 3.334 1.16 5.397zm9.213 10.214c-2.309 0-4.125-.897-5.45-2.69-1.312-1.793-1.968-4.277-1.968-7.453 0-3.223.668-5.73 2.004-7.524 1.347-1.804 3.199-2.706 5.554-2.706 2.473 0 4.36.96 5.66 2.882h.176c-.269-1.465-.404-2.771-.404-3.92V.648h5.379V28h-4.113l-1.038-2.549h-.228c-1.219 1.934-3.076 2.9-5.572 2.9zm1.88-4.272c1.372 0 2.374-.398 3.006-1.195.645-.797.996-2.15 1.055-4.06v-.58c0-2.11-.328-3.622-.984-4.536-.645-.914-1.7-1.371-3.164-1.371-1.196 0-2.127.51-2.795 1.53-.657 1.007-.985 2.478-.985 4.411 0 1.934.334 3.387 1.002 4.36.668.96 1.623 1.441 2.865 1.441z"
        fill="#2E2AC6"
      />
    </svg>
  ),

  'logo-md': (props) => (
    <svg width={139} height={31} fill="none" {...props}>
      <path
        d="M9.338 23.293c-1.924 0-3.433-.698-4.526-2.095h-.235c.156 1.367.235 2.158.235 2.373v6.636H.344V6.623h3.633l.63 2.124h.205c1.044-1.621 2.592-2.432 4.643-2.432 1.934 0 3.447.747 4.541 2.242 1.094 1.494 1.64 3.569 1.64 6.225 0 1.748-.258 3.267-.776 4.556-.508 1.289-1.235 2.27-2.182 2.944-.948.674-2.06 1.011-3.34 1.011zM8.02 9.89c-1.104 0-1.91.341-2.417 1.025-.508.674-.772 1.792-.791 3.355v.483c0 1.758.258 3.017.776 3.78.527.76 1.357 1.142 2.49 1.142 2.002 0 3.003-1.65 3.003-4.951 0-1.612-.249-2.818-.747-3.619-.488-.81-1.26-1.215-2.314-1.215zM33.245 23h-4.467v-9.565c0-2.364-.88-3.545-2.637-3.545-1.25 0-2.153.425-2.71 1.274-.557.85-.835 2.227-.835 4.13V23h-4.468V.207h4.468v4.644c0 .36-.034 1.21-.103 2.548l-.102 1.319h.234c.996-1.602 2.578-2.403 4.746-2.403 1.924 0 3.384.518 4.38 1.553s1.494 2.52 1.494 4.453V23zM43.427 9.494c-.947 0-1.69.303-2.226.908-.538.596-.845 1.446-.923 2.55h6.27c-.02-1.104-.308-1.954-.865-2.55-.557-.605-1.309-.908-2.256-.908zm.63 13.799c-2.637 0-4.697-.728-6.182-2.183-1.484-1.455-2.226-3.515-2.226-6.181 0-2.744.683-4.864 2.05-6.358 1.378-1.504 3.277-2.256 5.699-2.256 2.314 0 4.116.66 5.405 1.978 1.29 1.318 1.934 3.14 1.934 5.464v2.168H40.175c.049 1.27.425 2.26 1.128 2.973.703.713 1.69 1.07 2.96 1.07.985 0 1.918-.103 2.797-.308a14.004 14.004 0 002.754-.981v3.457c-.781.39-1.616.678-2.505.864-.889.195-1.973.293-3.252.293zM68.287 23h-4.468v-9.565c0-1.182-.21-2.066-.63-2.652-.42-.595-1.089-.893-2.007-.893-1.25 0-2.153.42-2.71 1.26-.556.83-.835 2.211-.835 4.145V23H53.17V6.623h3.414l.6 2.095h.25a4.683 4.683 0 012.05-1.787c.879-.41 1.875-.616 2.988-.616 1.905 0 3.35.518 4.336 1.553.987 1.026 1.48 2.51 1.48 4.453V23zm6.959-8.218c0 1.621.264 2.847.79 3.677.538.83 1.407 1.245 2.608 1.245 1.192 0 2.046-.41 2.564-1.23.527-.83.79-2.06.79-3.692 0-1.62-.263-2.837-.79-3.647-.528-.81-1.392-1.216-2.593-1.216-1.191 0-2.05.405-2.578 1.216-.527.8-.791 2.016-.791 3.647zm11.323 0c0 2.666-.703 4.751-2.11 6.255-1.406 1.504-3.364 2.256-5.873 2.256-1.573 0-2.96-.342-4.16-1.025-1.202-.694-2.124-1.685-2.769-2.974-.644-1.29-.967-2.793-.967-4.512 0-2.676.698-4.756 2.095-6.24 1.396-1.484 3.36-2.227 5.889-2.227 1.572 0 2.959.342 4.16 1.026 1.201.683 2.124 1.665 2.768 2.944.645 1.28.967 2.778.967 4.497zm11.486 8.511c-1.924 0-3.433-.698-4.527-2.095h-.234c.156 1.367.234 2.158.234 2.373v6.636H89.06V6.623h3.633l.63 2.124h.205c1.045-1.621 2.593-2.432 4.644-2.432 1.933 0 3.447.747 4.541 2.242 1.094 1.494 1.641 3.569 1.641 6.225 0 1.748-.259 3.267-.777 4.556-.508 1.289-1.235 2.27-2.182 2.944-.948.674-2.061 1.011-3.34 1.011zM96.736 9.89c-1.103 0-1.909.341-2.417 1.025-.508.674-.771 1.792-.79 3.355v.483c0 1.758.258 3.017.776 3.78.527.76 1.357 1.142 2.49 1.142 2.002 0 3.003-1.65 3.003-4.951 0-1.612-.25-2.818-.747-3.619-.489-.81-1.26-1.215-2.315-1.215zm13.668 4.892c0 1.621.264 2.847.792 3.677.537.83 1.406 1.245 2.607 1.245 1.191 0 2.046-.41 2.563-1.23.528-.83.791-2.06.791-3.692 0-1.62-.263-2.837-.791-3.647-.527-.81-1.391-1.216-2.592-1.216-1.192 0-2.051.405-2.578 1.216-.528.8-.792 2.016-.792 3.647zm11.324 0c0 2.666-.703 4.751-2.11 6.255-1.406 1.504-3.364 2.256-5.874 2.256-1.572 0-2.959-.342-4.16-1.025-1.201-.694-2.124-1.685-2.768-2.974-.645-1.29-.967-2.793-.967-4.512 0-2.676.698-4.756 2.095-6.24 1.396-1.484 3.359-2.227 5.888-2.227 1.572 0 2.959.342 4.16 1.026 1.202.683 2.124 1.665 2.769 2.944.644 1.28.967 2.778.967 4.497zm7.677 8.511c-1.924 0-3.438-.747-4.541-2.241-1.094-1.494-1.641-3.565-1.641-6.211 0-2.686.557-4.776 1.67-6.27 1.123-1.504 2.666-2.256 4.629-2.256 2.06 0 3.633.801 4.717 2.403h.146c-.224-1.22-.337-2.31-.337-3.267V.207h4.483V23h-3.428l-.864-2.124h-.191c-1.015 1.611-2.563 2.417-4.643 2.417zm1.567-3.56c1.143 0 1.978-.332 2.505-.996.537-.664.83-1.792.879-3.383v-.484c0-1.758-.274-3.018-.82-3.78-.538-.76-1.416-1.142-2.637-1.142-.996 0-1.773.425-2.329 1.275-.547.84-.821 2.065-.821 3.676 0 1.612.279 2.823.835 3.633.557.801 1.353 1.201 2.388 1.201z"
        fill="#2E2AC6"
      />
    </svg>
  ),

  'minus-outline': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 20a10 10 0 110-20 10 10 0 010 20zm0-2a8 8 0 100-16 8 8 0 000 16zm5-9v2H5V9h10z" />
    </svg>
  ),

  pause: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M5 4h3v12H5V4zm7 0h3v12h-3V4z" />
    </svg>
  ),

  phenopod: (props) => (
    <svg width={43} height={43} fill="none" {...props}>
      <path
        d="M33.012 37.826c.444.633 1.32.79 1.923.306a21 21 0 00-.462-33.084c-.616-.468-1.488-.286-1.913.36-.426.645-.244 1.51.367 1.983a18.2 18.2 0 01.397 28.442c-.597.49-.755 1.36-.312 1.993z"
        fill="#5A67D8"
      />
      <path
        d="M29.8 33.239c.444.633 1.322.792 1.913.293a15.4 15.4 0 00-.333-23.794c-.604-.483-1.477-.3-1.903.346-.426.646-.242 1.509.353 2.003a12.599 12.599 0 01.267 19.139c-.58.511-.74 1.38-.297 2.013zM28.77 21.77a7 7 0 11-14 0 7 7 0 0114 0zm-11.2 0a4.2 4.2 0 108.4 0 4.2 4.2 0 00-8.4 0z"
        fill="#5A67D8"
      />
      <path
        d="M14.77 21.77h2.8v18.9a1.4 1.4 0 11-2.8 0v-18.9z"
        fill="#5A67D8"
      />
    </svg>
  ),

  play: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M4 4l12 6-12 6z" />
    </svg>
  ),

  'play-outline': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M2.93 17.07A10 10 0 1117.07 2.93 10 10 0 012.93 17.07zm12.73-1.41A8 8 0 104.34 4.34a8 8 0 0011.32 11.32zM7 6l8 4-8 4V6z" />
    </svg>
  ),

  search: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M12.9 14.32a8 8 0 111.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 108 2a6 6 0 000 12z" />
    </svg>
  ),

  'twitter-color': (props) => (
    <svg viewBox="0 0 410.155 410.155" {...props}>
      <path
        d="M403.632 74.18a162.414 162.414 0 01-28.28 9.537 88.177 88.177 0 0023.275-37.067c1.295-4.051-3.105-7.554-6.763-5.385a163.188 163.188 0 01-43.235 17.862 11.02 11.02 0 01-2.702.336c-2.766 0-5.455-1.027-7.57-2.891-16.156-14.239-36.935-22.081-58.508-22.081-9.335 0-18.76 1.455-28.014 4.325-28.672 8.893-50.795 32.544-57.736 61.724-2.604 10.945-3.309 21.9-2.097 32.56a3.166 3.166 0 01-.797 2.481 3.278 3.278 0 01-2.753 1.091c-62.762-5.831-119.358-36.068-159.363-85.14-2.04-2.503-5.952-2.196-7.578.593-7.834 13.44-11.974 28.812-11.974 44.454 0 23.972 9.631 46.563 26.36 63.032a79.24 79.24 0 01-20.169-7.808c-3.06-1.7-6.825.485-6.868 3.985-.438 35.612 20.412 67.3 51.646 81.569a79.567 79.567 0 01-16.786-1.399c-3.446-.658-6.341 2.611-5.271 5.952 10.138 31.651 37.39 54.981 70.002 60.278-27.066 18.169-58.585 27.753-91.39 27.753l-10.227-.006c-3.151 0-5.816 2.054-6.619 5.106-.791 3.006.666 6.177 3.353 7.74 36.966 21.513 79.131 32.883 121.955 32.883 37.485 0 72.549-7.439 104.219-22.109 29.033-13.449 54.689-32.674 76.255-57.141 20.09-22.792 35.8-49.103 46.692-78.201 10.383-27.737 15.871-57.333 15.871-85.589v-1.346c-.001-4.537 2.051-8.806 5.631-11.712a174.776 174.776 0 0035.16-38.591c2.573-3.849-1.485-8.673-5.719-6.795z"
        fill="#76a9ea"
      />
    </svg>
  ),

  'user-solid-circle': (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M10 20a10 10 0 110-20 10 10 0 010 20zM7 6v2a3 3 0 106 0V6a3 3 0 10-6 0zm-3.65 8.44a8 8 0 0013.3 0 15.94 15.94 0 00-13.3 0z" />
    </svg>
  ),

  volume: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M5 7H1v6h4l5 5V2L5 7zm11.36 9.36l-1.41-1.41a6.98 6.98 0 000-9.9l1.41-1.41a8.97 8.97 0 010 12.72zm-2.82-2.82l-1.42-1.42a3 3 0 000-4.24l1.42-1.42a4.98 4.98 0 010 7.08z" />
    </svg>
  ),

  walk: (props) => (
    <svg viewBox="0 0 20 20" {...props}>
      <path d="M11 7l1.44 2.16c.31.47 1.01.84 1.57.84H17V8h-3l-1.44-2.16a5.94 5.94 0 00-1.4-1.4l-1.32-.88a1.72 1.72 0 00-1.7-.04L4 6v5h2V7l2-1-3 14h2l2.35-7.65L11 14v6h2v-8l-2.7-2.7L11 7zm1-3a2 2 0 100-4 2 2 0 000 4z" />
    </svg>
  ),
}
