<script lang="ts">
    import '../app.css'
    import '../js/declares'
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    import DataTable from '../components/DataTable.svelte'
    import { onMount } from 'svelte'

    export let siteName: string
    
    let newUsername = '';
    let errorMessage = '';
    let dataTableInstance: any;
    const tableId = 'users-table';

    // 定義用戶角色
    const roles = {
        ADMIN: '管理員',
        EDITOR: '編輯',
        USER: '一般用戶'
    };

    // 用戶數據
    let users = [
        {
            id: "1",
            username: "admin",
            role: "ADMIN",
            status: "active",
            created_at: "2024-03-15T08:00:00Z"
        },
        {
            id: "2",
            username: "editor01",
            role: "EDITOR",
            status: "active",
            created_at: "2024-03-16T10:30:00Z"
        },
        {
            id: "3",
            username: "writer_alice",
            role: "EDITOR",
            status: "suspended",
            created_at: "2024-03-17T14:15:00Z"
        },
        {
            id: "4",
            username: "moderator_bob",
            role: "EDITOR",
            status: "active",
            created_at: "2024-03-18T09:45:00Z"
        },
        {
            id: "5",
            username: "user_carol",
            role: "USER",
            status: "active",
            created_at: "2024-03-19T16:20:00Z"
        },
        {
            id: "6",
            username: "david_writer",
            role: "EDITOR",
            status: "active",
            created_at: "2024-03-20T11:25:00Z"
        },
        {
            id: "7",
            username: "emma_user",
            role: "USER",
            status: "suspended",
            created_at: "2024-03-21T13:40:00Z"
        },
        {
            id: "8",
            username: "frank_admin",
            role: "ADMIN",
            status: "active",
            created_at: "2024-03-22T15:55:00Z"
        },
        {
            id: "9",
            username: "grace_editor",
            role: "EDITOR",
            status: "active",
            created_at: "2024-03-23T08:10:00Z"
        },
        {
            id: "10",
            username: "henry_user",
            role: "USER",
            status: "active",
            created_at: "2024-03-24T10:05:00Z"
        }
    ];

    // API 處理函數 (與後端串接)
    const apiCreateUser = async (username: string) => {
        // TODO: 調用 API 創建用戶
        console.log('API - Creating user:', username);
        return true;
    };

    const apiUpdateUserRole = async (userId: string, newRole: string) => {
        // TODO: 調用 API 更新用戶角色
        console.log('API - Updating user role:', userId, newRole);
        return true;
    };

    const apiUpdateUserStatus = async (userId: string, newStatus: string) => {
        // TODO: 調用 API 更新用戶狀態
        console.log('API - Updating user status:', userId, newStatus);
        return true;
    };

    const apiDeleteUser = async (userId: string) => {
        // TODO: 調用 API 刪除用戶
        console.log('API - Deleting user:', userId);
        return true;
    };

    // UI 處理函數
    const uiHandleAddUser = async () => {
        if (!newUsername.trim()) {
            alert('請輸入用戶名');
            return;
        }

        try {
            const success = await apiCreateUser(newUsername);
            if (success) {
                const newUser = {
                    id: (users.length + 1).toString(),
                    username: newUsername,
                    role: 'USER',
                    status: 'active',
                    created_at: new Date().toISOString()
                };

                users = [...users, newUser];
                updateTable();
                alert('用戶新增成功');
                newUsername = '';
            }
        } catch (error) {
            alert('新增用戶失敗');
            console.error(error);
        }
    };

    const uiHandleEditRole = async (userId: string) => {
        const user = users.find(u => u.id === userId);
        if (!user) return;

        const newRole = prompt('請選擇新角色 (ADMIN/EDITOR/USER):', user.role);
        if (!newRole || !roles[newRole.toUpperCase()]) {
            alert('無效的角色！');
            return;
        }

        try {
            const success = await apiUpdateUserRole(userId, newRole.toUpperCase());
            if (success) {
                users = users.map(u => 
                    u.id === userId 
                        ? {...u, role: newRole.toUpperCase()}
                        : u
                );
                updateTable();
                alert('用戶角色更新成功');
            }
        } catch (error) {
            alert('更新角色失敗');
            console.error(error);
        }
    };

    const uiHandleToggleStatus = async (userId: string) => {
        const user = users.find(u => u.id === userId);
        if (!user) return;

        const newStatus = user.status === 'active' ? 'suspended' : 'active';
        const confirmMessage = newStatus === 'suspended' ? 
            '確定要停權此用戶嗎？' : 
            '確定要解除此用戶的停權狀態嗎？';

        if (!confirm(confirmMessage)) return;

        try {
            const success = await apiUpdateUserStatus(userId, newStatus);
            if (success) {
                users = users.map(u => 
                    u.id === userId 
                        ? {...u, status: newStatus}
                        : u
                );
                updateTable();
                alert(`用戶已${newStatus === 'active' ? '解除停權' : '停權'}`);
            }
        } catch (error) {
            alert('更新狀態失敗');
            console.error(error);
        }
    };

    const uiHandleDeleteUser = async (userId: string) => {
        if (!confirm('確定要刪除此用戶嗎？')) return;

        try {
            const success = await apiDeleteUser(userId);
            if (success) {
                users = users.filter(user => user.id !== userId);
                updateTable();
                alert('用戶刪除成功');
            }
        } catch (error) {
            alert('刪除用戶失敗');
            console.error(error);
        }
    };

    // 表格事件處理
    const uiHandleTableAction = (e: Event) => {
        const target = (e.target as HTMLElement).closest('button');
        if (!target) return;

        e.preventDefault();
        const userId = target.dataset.id;
        if (!userId) return;

        if (target.classList.contains('edit-role-btn')) {
            uiHandleEditRole(userId);
        } else if (target.classList.contains('toggle-status-btn')) {
            uiHandleToggleStatus(userId);
        } else if (target.classList.contains('delete-user-btn')) {
            uiHandleDeleteUser(userId);
        }
    };

    // 更新表格的工具函數
    const updateTable = () => {
        if (!dataTableInstance) return;
        
        try {
            dataTableInstance.clear();
            dataTableInstance.rows.add(users);
            dataTableInstance.draw();
            console.log('Table updated successfully');
        } catch (error) {
            console.error('Error updating table:', error);
        }
    };

    // 表格配置
    $: tableConfig = {
        data: users,
        columns: [
            { 
                data: 'id', 
                title: 'ID',
                width: '8%'
            },
            { 
                data: 'username', 
                title: '用戶名',
                width: '25%'
            },
            {
                data: 'role',
                title: '角色',
                width: '15%',
                render: (data) => `<span class="role-badge ${data.toLowerCase()}">${roles[data]}</span>`
            },
            {
                data: 'status',
                title: '狀態',
                width: '12%',
                render: (data) => `
                    <span class="status-badge ${data}">
                        ${data === 'active' ? '正常' : '已停權'}
                    </span>
                `
            },
            { 
                data: 'created_at', 
                title: '建立時間',
                width: '20%',
                render: (data) => new Date(data).toLocaleString()
            },
            {
                data: null,
                title: '操作',
                width: '20%',
                orderable: false,
                className: 'text-center',
                render: (data) => `
                    <div class="flex justify-center gap-2">
                        <button class="edit-role-btn" data-id="${data.id}" title="編輯角色">
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </button>
                        <button class="toggle-status-btn ${data.status === 'active' ? 'suspend' : 'activate'}" 
                                data-id="${data.id}" 
                                data-status="${data.status}"
                                title="${data.status === 'active' ? '停權用戶' : '解除停權'}">
                            ${data.status === 'active' ? 
                                '<svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" /></svg>' :
                                '<svg class="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 11V7a4 4 0 118 0m-4 8v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2z" /></svg>'
                            }
                        </button>
                        <button class="delete-user-btn" data-id="${data.id}" title="刪除用戶">
                            <svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                            </svg>
                        </button>
                    </div>
                `
            }
        ],
        searching: true,
        ordering: true,
        paging: true,
        pageLength: 10,
        lengthMenu: [5, 10, 25, 50],
        language: {
            "lengthMenu": "每頁 _MENU_ 筆",
            "zeroRecords": "沒有找到符合的記錄",
            "info": "顯示第 _START_ 至 _END_ 筆記錄，共 _TOTAL_ 筆",
            "infoEmpty": "顯示第 0 至 0 筆記錄，共 0 筆",
            "infoFiltered": "(從 _MAX_ 筆記錄中過濾)",
            "search": "搜尋：",
            "paginate": {
                "first": "首頁",
                "last": "末頁",
                "next": "下一頁",
                "previous": "上一頁"
            }
        },
        drawCallback: function(settings) {
            const table = document.getElementById(tableId);
            if (!table) return;
            
            // 移除所有現有的事件監聽器
            table.removeEventListener('click', uiHandleTableAction);
            
            // 添加新的事件監聽器
            table.addEventListener('click', uiHandleTableAction);
            
            console.log('Table redrawn, events reattached');
        }
    };

    onMount(() => {
        // 監聽 DataTable 實例創建完成的事件
        window.addEventListener('datatableCreated', (e: CustomEvent) => {
            if (e.detail.id === tableId) {
                dataTableInstance = e.detail.instance;
            }
        });
    });
</script>

<div class="min-h-screen flex flex-col bg-gray-50">
    <AdminNavbar {siteName} />
    
    <main class="flex-grow container mx-auto px-4 py-8">
        <div class="mb-8">
            <h1 class="text-3xl font-bold text-gray-800">用戶管理</h1>
            <p class="text-gray-600 mt-2">管理系統用戶、權限和狀態</p>
        </div>

        <!-- 新增用戶表單 -->
        <div class="bg-white p-6 rounded-lg shadow-md mb-6">
            <div class="mb-4">
                <h2 class="title is-4 mb-2">新增用戶</h2>
            </div>
            
            <div class="field is-grouped">
                <div class="control is-expanded has-icons-left">
                    <input
                        type="text"
                        bind:value={newUsername}
                        placeholder="輸入用戶名稱"
                        class="input"
                    />
                    <span class="icon is-small is-left">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                        </svg>
                    </span>
                </div>
                <div class="control">
                    <button
                        on:click={uiHandleAddUser}
                        class="button is-primary is-dark"
                        disabled={!newUsername.trim()}
                    >
                        <span class="icon is-small">
                            <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                            </svg>
                        </span>
                        <span>新增用戶</span>
                    </button>
                </div>
            </div>
            {#if errorMessage}
                <p class="help is-danger mt-2 is-flex is-align-items-center has-text-weight-medium is-size-6">
                    <span class="icon is-small mr-1">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                    </span>
                    {errorMessage}
                </p>
            {/if}
            <p class="help has-text-grey mt-3 has-text-weight-medium is-size-6">
                * 新增的用戶預設為一般用戶權限
            </p>
        </div>

        <!-- 用戶列表 -->
        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-6 text-gray-800">用戶列表</h2>
            <div class="overflow-hidden">
                <table id={tableId} class="display w-full">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>用戶名</th>
                            <th>角色</th>
                            <th>狀態</th>
                            <th>建立時間</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                </table>
                <DataTable id={tableId} config={tableConfig} />
            </div>
        </div>
    </main>

    <Footer {siteName} />
</div>

<style>
    /* 角色標籤樣式 */
    :global(.role-badge) {
        padding: 0.25rem 0.75rem;
        border-radius: 9999px;
        font-size: 0.875rem;
        font-weight: 500;
    }

    :global(.role-badge.admin) {
        background-color: #818cf8;
        color: white;
    }

    :global(.role-badge.editor) {
        background-color: #34d399;
        color: white;
    }

    :global(.role-badge.user) {
        background-color: #94a3b8;
        color: white;
    }

    /* 狀態標籤樣式 */
    :global(.status-badge) {
        padding: 0.25rem 0.75rem;
        border-radius: 9999px;
        font-size: 0.875rem;
        font-weight: 500;
    }

    :global(.status-badge.active) {
        background-color: #86efac;
        color: #166534;
    }

    :global(.status-badge.suspended) {
        background-color: #fca5a5;
        color: #991b1b;
    }

    /* 更新操作按鈕樣式 */
    :global(.edit-role-btn) {
        color: #3b82f6;
        padding: 0.5rem;
        border-radius: 0.375rem;
        transition: all 0.2s;
        border: none;
        background: transparent;
    }

    :global(.edit-role-btn:hover) {
        background-color: #eff6ff;
    }

    :global(.toggle-status-btn) {
        padding: 0.5rem;
        border-radius: 0.375rem;
        transition: all 0.2s;
        border: none;
        background: transparent;
    }

    :global(.toggle-status-btn.suspend:hover) {
        background-color: #fee2e2;
    }

    :global(.toggle-status-btn.activate:hover) {
        background-color: #ecfdf5;
    }

    :global(.delete-user-btn) {
        color: #ef4444;
        padding: 0.5rem;
        border-radius: 0.375rem;
        transition: all 0.2s;
        border: none;
        background: transparent;
    }

    :global(.delete-user-btn:hover) {
        background-color: #fee2e2;
    }

    /* 添加工具提示樣式 */
    :global([title]) {
        position: relative;
        cursor: pointer;
    }

    :global([title]:hover::after) {
        content: attr(title);
        position: absolute;
        bottom: 100%;
        left: 50%;
        transform: translateX(-50%);
        padding: 0.25rem 0.5rem;
        background-color: #1e293b;
        color: white;
        font-size: 0.75rem;
        border-radius: 0.25rem;
        white-space: nowrap;
        z-index: 10;
        margin-bottom: 0.25rem;
    }

    /* 搜尋框樣式 */
    :global(input[type="search"]) {
        background-color: white !important;
        border: 1px solid #e2e8f0 !important;
        border-radius: 0.375rem !important;
        padding: 0.5rem !important;
        margin-left: 0.5rem !important;
        outline: none !important;
        transition: all 0.2s !important;
    }

    :global(input[type="search"]:focus) {
        border-color: #3b82f6 !important;
        box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2) !important;
    }

    /* 調整搜尋框容器的樣式 */
    :global(.dataTables_wrapper .dataTables_filter) {
        margin-bottom: 1rem;
    }
</style>
