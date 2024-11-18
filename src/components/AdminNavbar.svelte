<script>
  import { onMount } from 'svelte';
  export let siteName;
  
  let isDropdownOpen = false;
  let isMobileMenuOpen = false;
  
  // 導航項目
  const navItems = [
    { label: '儀表板', icon: 'fa-chart-line', href: '/admin' },
    { label: '文章與類別', icon: 'fa-file-alt', href: '/admin' },
    { label: '用戶管理', icon: 'fa-users', href: '/admin/users' },
    { label: '設定', icon: 'fa-cog', href: '/admin/settings' }
  ];

  // 用戶選單項目
  const userMenuItems = [
    { label: '個人資料', icon: 'fa-user', href: '/admin/profile' },
    { label: '返回前台', icon: 'fa-home', href: '/' },
    { label: '登出', icon: 'fa-sign-out-alt', href: '/logout' }
  ];

  function toggleDropdown() {
    isDropdownOpen = !isDropdownOpen;
  }

  function toggleMobileMenu() {
    isMobileMenuOpen = !isMobileMenuOpen;
  }

  // 點擊外部關閉下拉選單
  function handleClickOutside(event) {
    const dropdown = document.querySelector('.user-dropdown');
    if (dropdown && !dropdown.contains(event.target)) {
      isDropdownOpen = false;
    }
  }

  onMount(() => {
    document.addEventListener('click', handleClickOutside);
    return () => {
      document.removeEventListener('click', handleClickOutside);
    };
  });
</script>

<nav class="admin-navbar">
  <div class="navbar-container">
    <!-- Logo 區域 -->
    <div class="navbar-brand">
      <a href="/admin" class="logo">
        <img src="/assets/logo.png" alt="Logo" class="logo-image">
        <span>{siteName} 後台</span>
      </a>
      <button class="mobile-toggle" on:click={toggleMobileMenu}>
        <i class="fas {isMobileMenuOpen ? 'fa-times' : 'fa-bars'}"></i>
      </button>
    </div>

    <!-- 主導航區域 -->
    <div class="navbar-main" class:active={isMobileMenuOpen}>
      <ul class="nav-items">
        {#each navItems as item}
          <li>
            <a href={item.href} class="nav-item">
              <i class="fas {item.icon}"></i>
              <span>{item.label}</span>
            </a>
          </li>
        {/each}
      </ul>
    </div>

    <!-- 用戶區域 -->
    <div class="navbar-user">
      <button class="user-button" on:click={toggleDropdown}>
        <img src="https://via.placeholder.com/32" alt="用戶頭像" class="user-avatar">
        <span class="user-name">管理員</span>
        <i class="fas fa-chevron-down"></i>
      </button>
      
      {#if isDropdownOpen}
        <div class="user-dropdown">
          {#each userMenuItems as item}
            <a href={item.href} class="dropdown-item">
              <i class="fas {item.icon}"></i>
              <span>{item.label}</span>
            </a>
          {/each}
        </div>
      {/if}
    </div>
  </div>
</nav>

<style>
  .admin-navbar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: 60px;
    background-color: white;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    z-index: 1000;
  }

  .navbar-container {
    max-width: 1400px;
    margin: 0 auto;
    height: 100%;
    display: flex;
    align-items: center;
    padding: 0 24px;
  }

  .navbar-brand {
    display: flex;
    align-items: center;
    margin-right: 48px;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 1.1em;
    font-weight: 600;
    color: var(--theme-text);
    text-decoration: none;
  }

  .logo-image {
    height: 24px;
    width: auto;
    display: block;
  }

  .navbar-main {
    flex: 1;
  }

  .nav-items {
    display: flex;
    gap: 8px;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 16px;
    color: var(--neutral-dark);
    text-decoration: none;
    border-radius: 6px;
    transition: all 0.3s ease;
  }

  .nav-item:hover {
    background-color: var(--neutral-light);
    color: var(--theme-accent);
  }

  .navbar-user {
    position: relative;
  }

  .user-button {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 6px 12px;
    border: none;
    background: none;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.3s ease;
  }

  .user-button:hover {
    background-color: var(--neutral-light);
  }

  .user-avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
  }

  .user-name {
    font-size: 0.95em;
    color: var(--theme-text);
  }

  .user-dropdown {
    position: absolute;
    top: 100%;
    right: 0;
    width: 200px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0,0,0,0.1);
    padding: 8px;
    margin-top: 8px;
  }

  .dropdown-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 10px 16px;
    color: var(--theme-text);
    text-decoration: none;
    border-radius: 6px;
    transition: all 0.3s ease;
  }

  .dropdown-item:hover {
    background-color: var(--neutral-light);
    color: var(--theme-accent);
  }

  .mobile-toggle {
    display: none;
    background: none;
    border: none;
    font-size: 1.2em;
    cursor: pointer;
    padding: 8px;
    color: var(--theme-text);
  }

  @media (max-width: 768px) {
    .mobile-toggle {
      display: block;
      margin-left: auto;
    }

    .navbar-main {
      position: fixed;
      top: 60px;
      left: 0;
      right: 0;
      background: white;
      padding: 16px;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
      display: none;
    }

    .navbar-main.active {
      display: block;
    }

    .nav-items {
      flex-direction: column;
    }

    .nav-item {
      padding: 12px 16px;
    }

    .navbar-user {
      margin-left: auto;
    }
  }
</style> 