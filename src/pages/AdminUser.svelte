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

<div class="container">
    <AdminNavbar {siteName} />
    
    <main class="section">
        <div class="content">
            <h1 class="title is-3">用戶管理</h1>
            <p class="subtitle is-size-8">管理系統用戶、權限和狀態</p>
        </div>

        <!-- 新增用戶表單 -->
        <div class="box mb-6">
            <div class="mb-4">
                <h2 class="title is-4">新增用戶</h2>
            </div>
            
            <div class="field is-grouped">
                <div class="control is-expanded">
                    <input
                        type="text"
                        bind:value={newUsername}
                        placeholder="輸入用戶名稱"
                        class="input"
                    />
                </div>
                <div class="control">
                    <button
                        on:click={uiHandleAddUser}
                        class="button is-primary"
                        disabled={!newUsername.trim()}
                    >
                        <span class="icon">
                            <i class="fas fa-plus"></i>
                        </span>
                        <span>新增用戶</span>
                    </button>
                </div>
            </div>
            
            {#if errorMessage}
                <p class="help is-danger">{errorMessage}</p>
            {/if}
            <p class="help is-size-6">* 新增的用戶預設為一般用戶權限</p>
        </div>

        <!-- 用戶列表 -->
        <div class="box">
            <h2 class="title is-4 mb-4">用戶列表</h2>
            <div class="table-container">
                <table id={tableId} class="table is-fullwidth is-striped">
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
    /* 更新樣式為 Bulma 兼容的樣式 */
    :global(.role-badge) {
        display: inline-block;
        padding: 0.25em 0.75em;
        border-radius: 290486px;
        font-size: 0.875em;
        font-weight: 500;
    }

    :global(.role-badge.admin) {
        background-color: hsl(217, 71%, 53%);
        color: white;
    }

    :global(.role-badge.editor) {
        background-color: hsl(141, 71%, 48%);
        color: white;
    }

    :global(.role-badge.user) {
        background-color: hsl(0, 0%, 71%);
        color: white;
    }

    :global(.status-badge) {
        display: inline-block;
        padding: 0.25em 0.75em;
        border-radius: 290486px;
        font-size: 0.875em;
        font-weight: 500;
    }

    :global(.status-badge.active) {
        background-color: hsl(141, 71%, 48%);
        color: white;
    }

    :global(.status-badge.suspended) {
        background-color: hsl(348, 100%, 61%);
        color: white;
    }

    :global(.button-group) {
        display: flex;
        gap: 0.5rem;
        justify-content: center;
    }

    :global(.button-group .button) {
        margin: 0;
    }
</style>
