export const useSidebar = () => {
  const open = useState('sidebar', () => true)
  return { open }
}
