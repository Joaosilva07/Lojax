import React, { useEffect, useState } from 'react';
import { DataGrid, GridToolbar } from '@mui/x-data-grid';
import { ListAllProdutos, CreateProduto, UpdateProduto, DeleteProduto } from "../wailsjs/go/bindings/ProdutoBindings";
import { ListarCategorias } from "../wailsjs/go/bindings/CategoriaBindings";
import { GetAllEstoques, CreateEstoque } from "../wailsjs/go/bindings/EstoqueBindings";
import AdminKitLayout from './AdminKitLayout';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import CategoryIcon from '@mui/icons-material/Category';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import OutlinedInput from '@mui/material/OutlinedInput';
import Checkbox from '@mui/material/Checkbox';
import ListItemText from '@mui/material/ListItemText';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import InputAdornment from '@mui/material/InputAdornment';
import SearchIcon from '@mui/icons-material/Search';
import AddBoxIcon from '@mui/icons-material/AddBox';
import SpeedDial from '@mui/material/SpeedDial';
import SpeedDialAction from '@mui/material/SpeedDialAction';

export default function Dashboard() {
  const [produtos, setProdutos] = useState([]);
  const [categorias, setCategorias] = useState([]);
  const [estoques, setEstoques] = useState([]);
  const [categoriaFiltro, setCategoriaFiltro] = useState('');
  const [estoqueFiltro, setEstoqueFiltro] = useState('todos');
  const [search, setSearch] = useState('');
  // Estados para modais
  const [openAddEdit, setOpenAddEdit] = useState(false);
  const [openDelete, setOpenDelete] = useState(false);
  const [openCategory, setOpenCategory] = useState(false);
  const [openStock, setOpenStock] = useState(false);
  const [selectedProduto, setSelectedProduto] = useState(null);
  // Estado do formulário de produto
  const [produtoForm, setProdutoForm] = useState({ nome_produto: '', classificacao: '', categorias: [] });
  // Estado do formulário de categorias
  const [categoriasSelecionadas, setCategoriasSelecionadas] = useState([]);
  // Estado do formulário de estoque
  const [estoqueForm, setEstoqueForm] = useState({ quantidade_em_estoque: 0 });
  // Modal de criar categoria
  const [openAddCategoria, setOpenAddCategoria] = useState(false);
  const [novaCategoria, setNovaCategoria] = useState('');

  useEffect(() => {
    // Testa se os bindings do Wails estão disponíveis
    const isWails = typeof window.go !== 'undefined' && window.go.bindings;
    if (isWails) {
      ListAllProdutos().then(setProdutos);
      ListarCategorias().then(setCategorias);
      GetAllEstoques().then(setEstoques);
    } else {
      // Fallback para desenvolvimento fora do Wails
      setProdutos([
        { id_produto: 1, nome_produto: "Produto Mock", classificacao: "A", categorias: [1] },
        { id_produto: 2, nome_produto: "Outro Produto", classificacao: "B", categorias: [2] }
      ]);
      setCategorias([
        { id: 1, nome_categoria: "Categoria Mock" },
        { id: 2, nome_categoria: "Outra Categoria" }
      ]);
      setEstoques([
        { id_produto: 1, quantidade_em_estoque: 10 },
        { id_produto: 2, quantidade_em_estoque: 0 }
      ]);
    }
  }, []);

  // Unifica os dados para a tabela
  const tabelaRows = produtos.map(produto => {
    // Pega o registro de estoque de maior id_estoque para o produto
    const estoquesProduto = estoques.filter(e => e.id_produto === produto.id_produto);
    const estoqueItem = estoquesProduto.length > 0 ? estoquesProduto.reduce((a, b) => (a.id_estoque > b.id_estoque ? a : b)) : null;
    const quantidade = estoqueItem ? estoqueItem.quantidade_em_estoque : 0;
    const status = quantidade > 0 ? 'Em Estoque' : 'Sem Estoque';
    const categoriaNomes = (produto.categorias || []).map(cid => {
      const cat = categorias.find(c => c.id === cid);
      return cat ? cat.nome_categoria : cid;
    });
    return {
      id: produto.id_produto,
      nome_produto: produto.nome_produto,
      classificacao: produto.classificacao,
      categorias: categoriaNomes,
      quantidade_em_estoque: quantidade,
      status,
      categorias_ids: produto.categorias || [],
    };
  });

  // Filtros
  const filteredRows = tabelaRows.filter(row => {
    // Filtro por categoria
    if (categoriaFiltro && !row.categorias.includes(categoriaFiltro)) return false;
    // Filtro por estoque
    if (estoqueFiltro === 'com' && row.quantidade_em_estoque <= 0) return false;
    if (estoqueFiltro === 'sem' && row.quantidade_em_estoque > 0) return false;
    // Filtro de busca
    if (search && !row.nome_produto.toLowerCase().includes(search.toLowerCase())) return false;
    return true;
  });

  // Colunas da tabela
  const columns = [
    { field: 'nome_produto', headerName: 'Produto', width: 180 },
    { field: 'classificacao', headerName: 'Classificação', width: 120 },
    { field: 'categorias', headerName: 'Categorias', width: 200, valueGetter: (params) => (params.row?.categorias ? params.row.categorias.join(', ') : '') },
    { field: 'quantidade_em_estoque', headerName: 'Quantidade', width: 120 },
    { field: 'status', headerName: 'Status', width: 120, renderCell: (params) => (
      <span style={{ color: params.value === 'Em Estoque' ? '#27ae60' : '#e74c3c', fontWeight: 700 }}>{params.value}</span>
    ) },
    {
      field: 'actions',
      headerName: 'Ações',
      width: 180,
      sortable: false,
      renderCell: (params) => (
        <div style={{ display: 'flex', gap: 8 }}>
          <IconButton size="small" color="primary" onClick={() => handleOpenEdit(params.row)}><EditIcon /></IconButton>
          <IconButton size="small" color="error" onClick={() => handleOpenDelete(params.row)}><DeleteIcon /></IconButton>
          <IconButton size="small" color="secondary" onClick={() => handleOpenCategory(params.row)}><CategoryIcon /></IconButton>
          <IconButton size="small" style={{ color: '#6d4caf' }} onClick={() => handleOpenStock(params.row)}><WarehouseIcon /></IconButton>
        </div>
      ),
    },
  ];

  const handleOpenAdd = () => {
    setSelectedProduto(null);
    setOpenAddEdit(true);
  };
  const handleOpenEdit = (produto) => {
    setSelectedProduto(produto);
    setOpenAddEdit(true);
  };
  const handleOpenDelete = (produto) => {
    setSelectedProduto(produto);
    setOpenDelete(true);
  };
  const handleOpenCategory = (produto) => {
    setSelectedProduto(produto);
    setOpenCategory(true);
  };
  const handleOpenStock = (produto) => {
    setSelectedProduto(produto);
    setOpenStock(true);
  };

  // Atualiza o formulário ao abrir para editar
  React.useEffect(() => {
    if (openAddEdit && selectedProduto) {
      setProdutoForm({
        nome_produto: selectedProduto.nome_produto || '',
        classificacao: selectedProduto.classificacao || '',
        categorias: selectedProduto.categorias || [],
      });
    } else if (openAddEdit) {
      setProdutoForm({ nome_produto: '', classificacao: '', categorias: [] });
    }
  }, [openAddEdit, selectedProduto]);

  // Atualiza categorias selecionadas ao abrir modal
  React.useEffect(() => {
    if (openCategory && selectedProduto) {
      setCategoriasSelecionadas(selectedProduto.categorias || []);
    }
  }, [openCategory, selectedProduto]);

  // Atualiza estoque ao abrir modal
  React.useEffect(() => {
    if (openStock && selectedProduto) {
      const estoqueItem = estoques.find(e => e.id_produto === selectedProduto.id);
      setEstoqueForm({ quantidade_em_estoque: estoqueItem ? estoqueItem.quantidade_em_estoque : 0 });
    }
  }, [openStock, selectedProduto, estoques]);

  // Handlers do formulário
  const handleProdutoFormChange = (e) => {
    const { name, value } = e.target;
    setProdutoForm((prev) => ({ ...prev, [name]: value }));
  };
  const handleProdutoCategoriasChange = (e) => {
    setProdutoForm((prev) => ({ ...prev, categorias: e.target.value }));
  };

  const handleProdutoSubmit = async () => {
    try {
      const categoriasIds = produtoForm.categorias.map(Number);
      if (selectedProduto) {
        await UpdateProduto({
          id_produto: selectedProduto.id,
          nome_produto: produtoForm.nome_produto,
          classificacao: produtoForm.classificacao,
          categorias: categoriasIds,
          descricao: selectedProduto.descricao || "",
          preco: selectedProduto.preco || 0,
        });
      } else {
        await CreateProduto({
          nome_produto: produtoForm.nome_produto,
          classificacao: produtoForm.classificacao,
          categorias: categoriasIds,
          descricao: "",
          preco: 0,
        });
      }
      setOpenAddEdit(false);
      const produtosAtualizados = await ListAllProdutos();
      setProdutos(produtosAtualizados);
    } catch (err) {
      alert("Erro ao salvar produto!");
    }
  };

  const handleCategoriasChange = (e) => {
    setCategoriasSelecionadas(e.target.value);
  };

  const handleSalvarCategorias = async () => {
    if (!selectedProduto) return;
    try {
      await UpdateProduto({
        ...selectedProduto,
        categorias: categoriasSelecionadas.map(Number),
      });
      setOpenCategory(false);
      const produtosAtualizados = await ListAllProdutos();
      setProdutos(produtosAtualizados);
    } catch (err) {
      alert('Erro ao atualizar categorias!');
    }
  };

  const handleEstoqueChange = (e) => {
    setEstoqueForm({ quantidade_em_estoque: Number(e.target.value) });
  };

  const handleSalvarEstoque = async () => {
    if (!selectedProduto) return;
    try {
      const produtoId = selectedProduto.id_produto || selectedProduto.id;
      const payload = {
        id_produto: produtoId,
        quantidade_em_estoque: Number(estoqueForm.quantidade_em_estoque),
      };
      console.log('Payload estoque:', payload);
      await CreateEstoque(payload);
      setOpenStock(false);
      const estoquesAtualizados = await GetAllEstoques();
      setEstoques(estoquesAtualizados);
      const produtosAtualizados = await ListAllProdutos();
      setProdutos(produtosAtualizados);
    } catch (err) {
      alert('Erro ao atualizar estoque!');
    }
  };

  const handleDeleteProduto = async () => {
    if (!selectedProduto) return;
    try {
      await DeleteProduto(selectedProduto.id);
      setOpenDelete(false);
      // Atualiza a lista
      const produtosAtualizados = await ListAllProdutos();
      setProdutos(produtosAtualizados);
    } catch (err) {
      alert('Erro ao deletar produto!');
    }
  };

  // Modal de criar categoria
  const handleAddCategoria = async () => {
    if (!novaCategoria.trim()) return;
    try {
      await window.go.bindings.CategoriaBindings.CreateCategoria({ nome_categoria: novaCategoria });
      setNovaCategoria('');
      setOpenAddCategoria(false);
      // Atualiza categorias
      const categoriasAtualizadas = await ListarCategorias();
      setCategorias(categoriasAtualizadas);
    } catch (err) {
      alert('Erro ao criar categoria!');
    }
  };

  return (
    <AdminKitLayout>
      <Box sx={{ width: '100%', px: { xs: 1, sm: 3 }, minHeight: '100vh', position: 'relative' }}>
        <Box sx={{
          display: 'flex', alignItems: 'center', justifyContent: 'space-between',
          pt: 3, pb: 1, width: '100%', position: 'sticky', top: 0, zIndex: 10, bgcolor: '#f4f6fa'
        }}>
          <Typography variant="h5" sx={{ fontWeight: 700, color: '#232e3c' }}>
            Produtos, Categorias e Estoque
          </Typography>
          <SpeedDial
            ariaLabel="Ações rápidas"
            icon={<AddBoxIcon fontSize="large" />}
            direction="left"
            sx={{ position: 'absolute', top: 16, right: 24 }}
          >
            <SpeedDialAction
              icon={<AddBoxIcon color="primary" />}
              tooltipTitle="Adicionar Produto"
              onClick={handleOpenAdd}
            />
            <SpeedDialAction
              icon={<CategoryIcon color="primary" />}
              tooltipTitle="Adicionar Categoria"
              onClick={() => setOpenAddCategoria(true)}
            />
          </SpeedDial>
        </Box>
        <Box sx={{
          display: 'flex', gap: 2, pb: 2, alignItems: 'center', flexWrap: 'wrap', width: '100%', flexDirection: { xs: 'column', sm: 'row' }
        }}>
          <TextField
            variant="outlined"
            size="small"
            placeholder="Buscar produto..."
            value={search}
            onChange={e => setSearch(e.target.value)}
            InputProps={{
              startAdornment: <InputAdornment position="start"><SearchIcon /></InputAdornment>
            }}
            sx={{ minWidth: 220, maxWidth: 320, flex: 1, width: { xs: '100%', sm: 'auto' } }}
          />
          <FormControl size="small" sx={{ minWidth: 180, width: { xs: '100%', sm: 'auto' } }}>
            <InputLabel>Categoria</InputLabel>
            <Select value={categoriaFiltro} onChange={e => setCategoriaFiltro(e.target.value)} label="Categoria">
              <MenuItem value="">Todas as categorias</MenuItem>
              {categorias.map(cat => (
                <MenuItem key={cat.id} value={cat.nome_categoria}>{cat.nome_categoria}</MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl size="small" sx={{ minWidth: 160, width: { xs: '100%', sm: 'auto' } }}>
            <InputLabel>Estoque</InputLabel>
            <Select value={estoqueFiltro} onChange={e => setEstoqueFiltro(e.target.value)} label="Estoque">
              <MenuItem value="todos">Todos Estoques</MenuItem>
              <MenuItem value="com">Com Estoque</MenuItem>
              <MenuItem value="sem">Sem Estoque</MenuItem>
            </Select>
          </FormControl>
        </Box>
        <Box sx={{ width: '100%', height: 'calc(100vh - 210px)', overflowX: 'auto' }}>
          <DataGrid
            rows={filteredRows}
            columns={columns}
            pageSize={10}
            sx={{
              border: 'none',
              fontFamily: 'Roboto, Inter, sans-serif',
              '& .MuiDataGrid-columnHeaders': { fontWeight: 700, fontSize: 16, bgcolor: '#f4f6fa' },
              '& .MuiDataGrid-cell': { fontSize: 15 },
              '& .MuiDataGrid-row:hover': { bgcolor: '#f5f5f5' },
              '& .MuiDataGrid-footerContainer': { borderTop: '1px solid #eee' }
            }}
          />
        </Box>
      </Box>
      {/* Modais */}
      <Dialog open={openAddEdit} onClose={() => setOpenAddEdit(false)} maxWidth="sm" fullWidth>
        <DialogTitle>{selectedProduto ? 'Editar Produto' : 'Adicionar Produto'}</DialogTitle>
        <DialogContent>
          <Box component="form" sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 1 }}>
            <TextField
              label="Nome do Produto"
              name="nome_produto"
              value={produtoForm.nome_produto}
              onChange={handleProdutoFormChange}
              fullWidth
              required
            />
            <TextField
              label="Classificação"
              name="classificacao"
              value={produtoForm.classificacao}
              onChange={handleProdutoFormChange}
              fullWidth
              required
            />
            <FormControl fullWidth>
              <InputLabel id="categorias-label">Categorias</InputLabel>
              <Select
                labelId="categorias-label"
                multiple
                name="categorias"
                value={produtoForm.categorias}
                onChange={handleProdutoCategoriasChange}
                input={<OutlinedInput label="Categorias" />}
                renderValue={(selected) =>
                  categorias
                    .filter((cat) => selected.includes(cat.id))
                    .map((cat) => cat.nome_categoria)
                    .join(', ')
                }
              >
                {categorias.map((cat) => (
                  <MenuItem key={cat.id} value={cat.id}>
                    <Checkbox checked={produtoForm.categorias.includes(cat.id)} />
                    <ListItemText primary={cat.nome_categoria} />
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenAddEdit(false)}>Cancelar</Button>
          <Button variant="contained" color="primary" onClick={handleProdutoSubmit}>Salvar</Button>
        </DialogActions>
      </Dialog>
      <Dialog open={openDelete} onClose={() => setOpenDelete(false)} maxWidth="xs">
        <DialogTitle>Confirmar exclusão</DialogTitle>
        <DialogContent>
          Tem certeza que deseja deletar o produto "{selectedProduto?.nome_produto}"?
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenDelete(false)}>Cancelar</Button>
          <Button variant="contained" color="error" onClick={handleDeleteProduto}>Deletar</Button>
        </DialogActions>
      </Dialog>
      <Dialog open={openCategory} onClose={() => setOpenCategory(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Gerenciar Categorias</DialogTitle>
        <DialogContent>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel id="categorias-produto-label">Categorias</InputLabel>
            <Select
              labelId="categorias-produto-label"
              multiple
              value={categoriasSelecionadas}
              onChange={handleCategoriasChange}
              input={<OutlinedInput label="Categorias" />}
              renderValue={(selected) =>
                categorias
                  .filter((cat) => selected.includes(cat.id))
                  .map((cat) => cat.nome_categoria)
                  .join(', ')
              }
            >
              {categorias.map((cat) => (
                <MenuItem key={cat.id} value={cat.id}>
                  <Checkbox checked={categoriasSelecionadas.includes(cat.id)} />
                  <ListItemText primary={cat.nome_categoria} />
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenCategory(false)}>Fechar</Button>
          <Button variant="contained" color="primary" onClick={handleSalvarCategorias}>Salvar</Button>
        </DialogActions>
      </Dialog>
      <Dialog open={openStock} onClose={() => setOpenStock(false)} maxWidth="xs">
        <DialogTitle>Atualizar Estoque</DialogTitle>
        <DialogContent>
          <TextField
            label="Quantidade em Estoque"
            type="number"
            value={estoqueForm.quantidade_em_estoque}
            onChange={handleEstoqueChange}
            fullWidth
            sx={{ mt: 2 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenStock(false)}>Cancelar</Button>
          <Button variant="contained" color="primary" onClick={handleSalvarEstoque}>Salvar</Button>
        </DialogActions>
      </Dialog>
      {/* Modal de criar categoria */}
      <Dialog open={openAddCategoria} onClose={() => setOpenAddCategoria(false)} maxWidth="xs">
        <DialogTitle>Adicionar Categoria</DialogTitle>
        <DialogContent>
          <TextField
            label="Nome da Categoria"
            value={novaCategoria}
            onChange={e => setNovaCategoria(e.target.value)}
            fullWidth
            autoFocus
            sx={{ mt: 2 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenAddCategoria(false)}>Cancelar</Button>
          <Button variant="contained" color="primary" onClick={handleAddCategoria}>Adicionar</Button>
        </DialogActions>
      </Dialog>
    </AdminKitLayout>
  );
} 